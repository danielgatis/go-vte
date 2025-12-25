package vte

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSosBellTerminated(t *testing.T) {
	// ESC X data BEL
	input := "\x1bXsome data\x07"
	expected := testSosPmApcSequence{
		kind:           SosKind,
		data:           []byte("some data"),
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

func TestSosEscBackslashTerminated(t *testing.T) {
	// ESC X data ESC \
	input := "\x1bXsome data\x1b\\"
	expected := []testSequence{
		testSosPmApcSequence{
			kind:           SosKind,
			data:           []byte("some data"),
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

func TestEmptySos(t *testing.T) {
	// ESC X BEL
	input := "\x1bX\x07"
	expected := testSosPmApcSequence{
		kind:           SosKind,
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

func TestSosTerminatedByCan(t *testing.T) {
	// ESC X data CAN (0x18)
	input := "\x1bXsome data\x18"
	expected := testSosPmApcSequence{
		kind:           SosKind,
		data:           []byte("some data"),
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
