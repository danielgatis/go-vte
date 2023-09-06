package vte

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscReset(t *testing.T) {
	expected := testEscSequence{
		intermediates: []byte{'('},
		ignore:        false,
		b:             'A',
	}
	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte("\x1b[3;1\x1b(A") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0].(testEscSequence))
}
