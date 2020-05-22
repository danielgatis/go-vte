package utf8

const continuationMask byte = 0b0011_1111

// State represents a state type
type State = byte

// Action represents a action type
type Action = byte

type codePointCallback func(char uint32)
type invalidSequenceCallback func()

// Parser represents a state machine
type Parser struct {
	codepoint uint32
	state     State
	ccb       codePointCallback
	icb       invalidSequenceCallback
}

// New returns a new parser
func New(ccb codePointCallback, icb invalidSequenceCallback) *Parser {
	return &Parser{
		codepoint: 0,
		state:     groundState,
		ccb:       ccb,
		icb:       icb,
	}
}

// Advance advances the state machine
func (p *Parser) Advance(b byte) {
	change := stateTable[p.state][b]
	state := change & 0x0f
	action := change >> 4

	p.performAction(action, b)
	p.state = state
}

// Codepoint returns a codepoint
func (p *Parser) Codepoint() uint32 {
	return p.codepoint
}

// StateName returns the current state name
func (p *Parser) StateName() string {
	return stateNames[p.state]
}

// State returns the current state
func (p *Parser) State() State {
	return p.state
}

func (p *Parser) performAction(action, b byte) {
	switch action {
	case invalidSequenceAction:
		p.codepoint = 0
		p.icb()

	case emitByteAction:
		p.ccb(uint32(b))

	case setByte1Action:
		point := p.codepoint | uint32(b&continuationMask)
		p.codepoint = 0
		p.ccb(point)

	case setByte2Action:
		p.codepoint = p.codepoint | uint32((b&continuationMask))<<6

	case setByte2TopAction:
		p.codepoint = p.codepoint | uint32((b&0b0001_1111))<<6

	case setByte3Action:
		p.codepoint = p.codepoint | uint32((b&continuationMask))<<12

	case setByte3TopAction:
		p.codepoint = p.codepoint | uint32((b&0b0000_1111))<<12

	case setByte4Action:
		p.codepoint = p.codepoint | uint32((b&0b0000_0111))<<18
	}
}
