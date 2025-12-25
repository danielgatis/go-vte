package vte

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPmBellTerminated(t *testing.T) {
	// ESC ^ data BEL
	input := "\x1b^private message\x07"
	expected := testSosPmApcSequence{
		kind:           PmKind,
		data:           []byte("private message"),
		bellTerminated: true,
	}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestPmEscBackslashTerminated(t *testing.T) {
	// ESC ^ data ESC \
	input := "\x1b^private message\x1b\\"
	expected := []testSequence{
		testSosPmApcSequence{
			kind:           PmKind,
			data:           []byte("private message"),
			bellTerminated: false,
		},
		testEscSequence{
			intermediates: []byte{},
			ignore:        false,
			b:             '\\',
		},
	}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 2, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}

func TestEmptyPm(t *testing.T) {
	// ESC ^ BEL
	input := "\x1b^\x07"
	expected := testSosPmApcSequence{
		kind:           PmKind,
		data:           []byte{},
		bellTerminated: true,
	}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestPmTerminatedByCan(t *testing.T) {
	// ESC ^ data CAN (0x18)
	input := "\x1b^private message\x18"
	expected := testSosPmApcSequence{
		kind:           PmKind,
		data:           []byte("private message"),
		bellTerminated: false,
	}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}
