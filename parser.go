package vte

import (
	"math"

	"github.com/danielgatis/go-vte/utf8"
)

// State represents a state type
type State = byte

// Action represents a action type
type Action = byte

// MaxIntermediates is the maximum number of intermediates allowed.
const MaxIntermediates = 2

// MaxOscRaw is the maximum number of bytes allowed in an osc parameter.
const MaxOscRaw = 1024

// MaxOscParams is the maximum number of osc parameters allowed.
const MaxOscParams = 16

// Performer is an interface for parsing.
type Performer interface {
	// Print is called when a print action is performed.
	Print(r rune)

	// Execute is called when an execute action is performed.
	Execute(b byte)

	// Put is called when a put action is performed.
	Put(b byte)

	// Unhook is called when an unhook action is performed.
	Unhook()

	// Hook is called when a hook action is performed.
	Hook(params *Params, intermediates []byte, ignore bool, r rune)

	// OscDispatch is called when an osc dispatch action is performed.
	OscDispatch(params [][]byte, bellTerminated bool)

	// CsiDispatch is called when a csi dispatch action is performed.
	CsiDispatch(params *Params, intermediates []byte, ignore bool, r rune)

	// EscDispatch is called when an esc dispatch action is performed.
	EscDispatch(intermediates []byte, ignore bool, b byte)
}

// Parser represents a state machine.
type Parser struct {
	state           byte
	intermediates   [MaxIntermediates]uint8
	intermediateIdx int
	params          *Params
	param           uint16
	oscRaw          []byte
	oscParams       [MaxOscParams][2]int
	oscNumParams    int
	ignoring        bool
	utf8Parser      *utf8.Parser

	performer Performer
}

func (p *Parser) prtcb(char rune) {
	if p.performer != nil {
		p.performer.Print(char)
	}
}

func (p *Parser) execb(b byte) {
	if p.performer != nil {
		p.performer.Execute(b)
	}
}

func (p *Parser) putcb(b byte) {
	if p.performer != nil {
		p.performer.Put(b)
	}
}

func (p *Parser) uhocb() {
	if p.performer != nil {
		p.performer.Unhook()
	}
}

func (p *Parser) hokcb(params *Params, intermediates []byte, ignore bool, r rune) {
	if p.performer != nil {
		p.performer.Hook(params, intermediates, ignore, r)
	}
}

func (p *Parser) osccb(params [][]byte, bellTerminated bool) {
	if p.performer != nil {
		p.performer.OscDispatch(params, bellTerminated)
	}
}

func (p *Parser) csicb(params *Params, intermediates []byte, ignore bool, r rune) {
	if p.performer != nil {
		p.performer.CsiDispatch(params, intermediates, ignore, r)
	}
}

func (p *Parser) esccb(intermediates []byte, ignore bool, b byte) {
	if p.performer != nil {
		p.performer.EscDispatch(intermediates, ignore, b)
	}
}

// New returns a new parser.
func New(
	performer Performer,
) *Parser {
	p := &Parser{
		params:    NewParams(),
		state:     GroundState,
		performer: performer,
	}

	p.utf8Parser = utf8.New(
		func(r rune) {
			p.prtcb(r)
			p.state = GroundState
		},

		func() {
			p.prtcb('�')
			p.state = GroundState
		},
	)

	return p
}

// Advance advances the state machine.
func (p *Parser) Advance(b byte) {
	if p.state == Utf8State {
		p.processUtf8(b)
	} else {
		change := stateTable[AnywhereState][b]

		if change == 0 {
			change = stateTable[p.state][b]
		}

		state := change & 0x0f
		action := change >> 4

		p.performStateChange(state, action, b)
	}
}

// Intermediates returns the intermediates
func (p *Parser) Intermediates() []byte {
	return append([]byte{}, p.intermediates[:p.intermediateIdx]...)
}

// Params returns the params
func (p *Parser) Params() *Params {
	return p.params
}

// OscParams returns the osc params
func (p *Parser) OscParams() [][]byte {
	params := make([][]byte, 0, MaxOscParams)

	for i := 0; i < p.oscNumParams; i++ {
		indices := p.oscParams[i]
		param := p.oscRaw[indices[0]:indices[1]]
		params = append(params, param)
	}

	return params[:p.oscNumParams]
}

// StateName returns the current state name
func (p *Parser) StateName() string {
	return stateNames[p.state]
}

// State returns the current state
func (p *Parser) State() State {
	return p.state
}

func (p *Parser) processUtf8(b byte) {
	p.utf8Parser.Advance(b)
}

func (p *Parser) performStateChange(state, action, b byte) {
	switch state {
	case AnywhereState:
		p.performAction(action, b)
	default:
		switch p.state {
		case DcsPassthroughState:
			p.performAction(UnhookAction, b)
		case OscStringState:
			p.performAction(OscEndAction, b)
		default:
			break
		}

		switch action {
		case NoneAction:
			break
		default:
			p.performAction(action, b)
		}

		switch state {
		case CsiEntryState, DcsEntryState, EscapeState:
			p.performAction(ClearAction, b)
		case DcsPassthroughState:
			p.performAction(HookAction, b)
		case OscStringState:
			p.performAction(OscStartAction, b)
		default:
			break
		}

		// Assume the new state
		p.state = state
	}
}

func (p *Parser) performAction(action, b byte) {
	switch action {
	case IgnoreAction:
		break

	case NoneAction:
		break

	case PrintAction:
		p.prtcb(rune(b))

	case ExecuteAction:
		p.execb(b)

	case HookAction:
		if p.params.IsFull() {
			p.ignoring = true
		} else {
			p.params.Push(p.param)
		}

		p.hokcb(
			p.Params(),
			p.Intermediates(),
			p.ignoring,
			rune(b),
		)

	case PutAction:
		p.putcb(b)

	case OscStartAction:
		p.oscRaw = make([]byte, 0)
		p.oscNumParams = 0

	case OscPutAction:
		idx := len(p.oscRaw)

		if b == ';' {
			paramIdx := p.oscNumParams
			switch paramIdx {
			case MaxOscParams:
				return
			case 0:
				p.oscParams[paramIdx] = [2]int{0, idx}
			default:
				prev := p.oscParams[paramIdx-1]
				begin := prev[1]
				p.oscParams[paramIdx] = [2]int{begin, idx}
			}
			p.oscNumParams++
		} else {
			p.oscRaw = append(p.oscRaw, b)
		}

	case OscEndAction:
		paramIdx := p.oscNumParams
		idx := len(p.oscRaw)

		switch paramIdx {
		case MaxOscParams:
			break
		case 0:
			p.oscParams[paramIdx] = [2]int{0, idx}
			p.oscNumParams++
		default:
			prev := p.oscParams[paramIdx-1]
			begin := prev[1]
			p.oscParams[paramIdx] = [2]int{begin, idx}
			p.oscNumParams++
		}

		p.osccb(
			p.OscParams(),
			b == 0x07,
		)

	case UnhookAction:
		p.uhocb()

	case CsiDispatchAction:
		if p.params.IsFull() {
			p.ignoring = true
		} else {
			p.params.Push(p.param)
		}

		p.csicb(
			p.Params(),
			p.Intermediates(),
			p.ignoring,
			rune(b),
		)

	case EscDispatchAction:
		p.esccb(
			p.Intermediates(),
			p.ignoring,
			b,
		)

	case CollectAction:
		if p.intermediateIdx == MaxIntermediates {
			p.ignoring = true
		} else {
			p.intermediates[p.intermediateIdx] = b
			p.intermediateIdx++
		}

	case ParamAction:
		if p.params.IsFull() {
			p.ignoring = true
			return
		}

		if b == ';' {
			p.params.Push(p.param)
			p.param = 0
		} else if b == ':' {
			p.params.Extend(p.param)
			p.param = 0
		} else {
			p.param = smulu16(p.param, 10)
			p.param = saddu16(p.param, uint16((b - '0')))
		}

	case ClearAction:
		// Reset everything on ESC/CSI/DCS entry
		p.intermediateIdx = 0
		p.ignoring = false
		p.param = 0

		p.params.Clear()

	case BeginUtf8Action:
		p.processUtf8(b)
	}
}

func saddu16(a, b uint16) uint16 {
	if b > 0 && a > math.MaxUint16-b {
		return math.MaxUint16
	}

	return a + b
}

func smulu16(a, b uint16) uint16 {
	if a > 0 && b > 0 && a > math.MaxUint16/b {
		return math.MaxUint16
	}

	return a * b
}
