package vte

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApcBellTerminated(t *testing.T) {
	// ESC _ Gf=100,s=10 BEL
	input := "\x1b_Gf=100,s=10\x07"
	expected := testSosPmApcSequence{
		kind:           ApcKind,
		data:           []byte("Gf=100,s=10"),
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

func TestApcEscBackslashTerminated(t *testing.T) {
	// ESC _ Gf=100,s=10 ESC \
	input := "\x1b_Gf=100,s=10\x1b\\"
	expected := []testSequence{
		testSosPmApcSequence{
			kind:           ApcKind,
			data:           []byte("Gf=100,s=10"),
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

func TestEmptyApc(t *testing.T) {
	// ESC _ BEL
	input := "\x1b_\x07"
	expected := testSosPmApcSequence{
		kind:           ApcKind,
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

func TestApcWithKittyGraphicsPayload(t *testing.T) {
	// Simulating Kitty Graphics Protocol: ESC _ Gf=100,s=10,v=20,m=1;base64... ESC \
	input := "\x1b_Gf=100,s=10,v=20,m=1;SGVsbG8gV29ybGQh\x1b\\"
	expected := []testSequence{
		testSosPmApcSequence{
			kind:           ApcKind,
			data:           []byte("Gf=100,s=10,v=20,m=1;SGVsbG8gV29ybGQh"),
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

func TestApcTerminatedByCan(t *testing.T) {
	// ESC _ data CAN (0x18)
	input := "\x1b_some data\x18"
	expected := testSosPmApcSequence{
		kind:           ApcKind,
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

func TestApcTerminatedBySub(t *testing.T) {
	// ESC _ data SUB (0x1a)
	input := "\x1b_some data\x1a"
	expected := testSosPmApcSequence{
		kind:           ApcKind,
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

func TestApcMaxBufferSize(t *testing.T) {
	// Test that data beyond maxSosPmApcRaw is truncated
	numBytes := maxSosPmApcRaw + 100
	inputStart := []byte{0x1b, '_'}
	inputEnd := []byte{0x07}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range inputStart {
		parser.Advance(b)
	}

	for i := 0; i < numBytes; i++ {
		parser.Advance('a')
	}

	for _, b := range inputEnd {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testSosPmApcSequence)
	assert.Equal(t, ApcKind, seq.kind)
	assert.Equal(t, maxSosPmApcRaw, len(seq.data))
	assert.Equal(t, bytes.Repeat([]byte{'a'}, maxSosPmApcRaw), seq.data)
	assert.Equal(t, true, seq.bellTerminated)
}

func TestMultipleApcSequences(t *testing.T) {
	// Two APC sequences in a row
	input := "\x1b_first\x07\x1b_second\x07"
	expected := []testSequence{
		testSosPmApcSequence{
			kind:           ApcKind,
			data:           []byte("first"),
			bellTerminated: true,
		},
		testSosPmApcSequence{
			kind:           ApcKind,
			data:           []byte("second"),
			bellTerminated: true,
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

func TestApcWithBinaryData(t *testing.T) {
	// APC with binary data (including high bytes)
	input := []byte{0x1b, '_', 0x20, 0x7e, 0x80, 0xff, 0x07}
	expected := testSosPmApcSequence{
		kind:           ApcKind,
		data:           []byte{0x20, 0x7e, 0x80, 0xff},
		bellTerminated: true,
	}

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range input {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}
