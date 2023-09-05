package vte

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var oscBytes = []byte{
	0x1b, 0x5d, // Begin OSC
	'2', ';', 'j', 'w', 'i', 'l', 'm', '@', 'j', 'w', 'i', 'l', 'm', '-', 'd',
	'e', 's', 'k', ':', ' ', '~', '/', 'c', 'o', 'd', 'e', '/', 'a', 'l', 'a',
	'c', 'r', 'i', 't', 't', 'y', 0x07, // End OSC
}

func TestOsc(t *testing.T) {
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range oscBytes {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, 2, len(dispatcher.dispatched[0].(testOscSequence).params))
	assert.Equal(t, oscBytes[2:3], dispatcher.dispatched[0].(testOscSequence).params[0])
	assert.Equal(t, oscBytes[4:len(oscBytes)-1], dispatcher.dispatched[0].(testOscSequence).params[1])
	assert.Equal(t, true, dispatcher.dispatched[0].(testOscSequence).bell)
}

func TestEmptyOsc(t *testing.T) {
	expected := testOscSequence{
		params: [][]byte{{}},
		bell:   true,
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte{0x1b, 0x5d, 0x07} { // ESC ] BEL ("\x1b]\x07")
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestOscMaxParams(t *testing.T) {
	input := fmt.Sprintf("\x1b]%s\x1b", strings.Repeat(";", maxOscParams+1))
	expected := testOscSequence{
		params: [][]byte{},
	}
	for i := 0; i < maxOscParams; i++ {
		expected.params = append(expected.params, []byte{})
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestOscBellTerminated(t *testing.T) {
	input := "\x1b]11;ff/00/ff\x07"
	expected := testOscSequence{
		params: [][]byte{
			[]byte("11"), []byte("ff/00/ff"),
		},
		bell: true,
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestOscC0StTerminated(t *testing.T) {
	input := "\x1b]11;ff/00/ff\x1b\\"
	expected := []testSequence{
		testOscSequence{
			params: [][]byte{
				[]byte("11"), []byte("ff/00/ff"),
			},
			bell: false,
		},
		testEscSequence{
			intermediates: []byte{},
			ignore:        false,
			b:             '\\',
		},
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 2, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}

func TestOscWithUTF8Arguments(t *testing.T) {
	input := []byte{
		0x0d, 0x1b, 0x5d, 0x32, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x27, 0xc2, 0xaf, 0x5c,
		0x5f, 0x28, 0xe3, 0x83, 0x84, 0x29, 0x5f, 0x2f, 0xc2, 0xaf, 0x27, 0x20, 0x26, 0x26,
		0x20, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x20, 0x31, 0x07,
	}

	expected := testOscSequence{
		params: [][]byte{
			[]byte("2"), []byte(`echo '¯\_(ツ)_/¯' && sleep 1`),
		},
		bell: true,
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range input {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestOscContainingStringTerminator(t *testing.T) {
	input := []byte("\x1b]2;\xe6\x9c\xab\x1b\\")

	expected := []testSequence{
		testOscSequence{
			params: [][]byte{
				[]byte("2"), []byte("\xe6\x9c\xab"),
			},
			bell: false,
		},
		testEscSequence{
			intermediates: []byte{},
			ignore:        false,
			b:             '\\',
		},
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range input {
		parser.Advance(b)
	}

	assert.Equal(t, 2, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}

func TestOcsExceedMaxBufferSize(t *testing.T) {
	numBytes := maxOscRaw + 100
	inputStart := []byte{0x1b, ']', '5', '2', ';', 's'}
	inputEnd := []byte{0x07}

	expected := testOscSequence{
		params: [][]byte{
			[]byte("52"), append([]byte{'s'}, bytes.Repeat([]byte{'a'}, numBytes)...),
		},
		bell: true,
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

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
	assert.Equal(t, len(expected.params[1]), numBytes+len(inputEnd))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}
