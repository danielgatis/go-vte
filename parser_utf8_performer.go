package vte

import (
	"github.com/danielgatis/go-utf8"
)

var _ (utf8.Performer) = (*utf8Performer)(nil)

// utf8Performer is a Performer implementation for utf8.
type utf8Performer struct {
	prtcb    func(r rune)
	setState func(state State)
}

// newUtf8Performer returns a new utf8Performer.
func newUtf8Performer(prtcb func(r rune), setState func(state State)) utf8.Performer {
	return &utf8Performer{prtcb: prtcb, setState: setState}
}

// Execute is called when a valid utf8 sequence is received.
func (p *utf8Performer) CodePoint(r rune) {
	p.prtcb(r)
	p.setState(GroundState)
}

// InvalidSequece is called when an invalid utf8 sequence is received.
func (p *utf8Performer) InvalidSequece() {
	p.prtcb('�')
	p.setState(GroundState)
}
