package vte

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDcsMaxParams(t *testing.T) {
	input := fmt.Sprintf("\x1bP%sp", strings.Repeat("1;", maxParams+1))
	expected := testDcsHookSequence{
		params: [][]uint16{
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
		},
		ignore: true,
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected.ignore, dispatcher.dispatched[0].(testDcsHookSequence).ignore)
	assert.Equal(t, maxParams, len(dispatcher.dispatched[0].(testDcsHookSequence).params))
	assert.Equal(t, expected.params, dispatcher.dispatched[0].(testDcsHookSequence).params)
}

func TestDcsReset(t *testing.T) {
	input := "\x1b[3;1\x1bP1$tx\x9c"
	expected := []testSequence{
		testDcsHookSequence{
			params:        [][]uint16{{1}},
			intermediates: []byte{'$'},
			rune:          't',
			ignore:        false,
		},
		testDcsPutSequence('x'),
		testDcsUnhookSequence{},
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 3, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}

func TestDcsParse(t *testing.T) {
	input := "\x1bP0;1|17/ab\x9c"
	expected := []testSequence{
		testDcsHookSequence{
			params:        [][]uint16{{0}, {1}},
			rune:          '|',
			intermediates: []byte{},
			ignore:        false,
		},
		testDcsPutSequence('1'),
		testDcsPutSequence('7'),
		testDcsPutSequence('/'),
		testDcsPutSequence('a'),
		testDcsPutSequence('b'),
		testDcsUnhookSequence{},
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, len(expected), len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}

// FIXME: This test is broken.
func TestDcsIntermediateResetOnExit(t *testing.T) {
	input := "\x1bP=1sZZZ\x1b+\x5c"
	expected := []testSequence{
		testDcsHookSequence{
			params:        [][]uint16{{1}},
			intermediates: []byte{'='},
			rune:          's',
			ignore:        false,
		},
		testDcsPutSequence('Z'),
		testDcsPutSequence('Z'),
		testDcsPutSequence('Z'),
		testDcsUnhookSequence{},
		testEscSequence{
			intermediates: []byte{'+'},
			b:             0x5c,
			ignore:        false,
		},
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, len(expected), len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched)
}
