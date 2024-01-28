package vte

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockParser struct {
	lastRune rune
	state    State
}

func (p *MockParser) prtcb(r rune) {
	p.lastRune = r
}

func (p *MockParser) setState(state State) {
	p.state = state
}

func TestUtf8Performer(t *testing.T) {
	mockParser := &MockParser{}
	performer := newUtf8Performer(mockParser.prtcb, mockParser.setState)

	performer.CodePoint('A')
	assert.Equal(t, 'A', mockParser.lastRune)
	assert.Equal(t, GroundState, mockParser.state)

	performer.InvalidSequece()
	assert.Equal(t, '�', mockParser.lastRune)
	assert.Equal(t, GroundState, mockParser.state)
}
