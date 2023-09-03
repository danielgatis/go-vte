package vte

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsiMaxParams(t *testing.T) {
	expected := testCsiSequence{
		params: [][]uint16{
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {0},
		},
		ignore:        false,
		intermediates: []byte{},
		rune:          'p',
	}

	input := "\x1b[" + strings.Repeat("1;", maxParams-1) + "p"
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestCsiParamsIgnoreLong(t *testing.T) {
	expected := testCsiSequence{
		params: [][]uint16{
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
			{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1},
		},
		ignore:        true,
		intermediates: []byte{},
		rune:          'p',
	}
	strParams := "\x1b[" + strings.Repeat("1;", maxParams) + "p"

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(strParams) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestCsiParamsTrailingSemicolon(t *testing.T) {
	expected := testCsiSequence{
		params:        [][]uint16{{4}, {0}},
		intermediates: []byte{},
		ignore:        false,
		rune:          'm',
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[4;m") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestCsiParamsLeadingSemicolon(t *testing.T) {
	expected := testCsiSequence{
		params:        [][]uint16{{0}, {4}},
		intermediates: []byte{},
		ignore:        false,
		rune:          'm',
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[;4m") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestLongCsiParam(t *testing.T) {
	expected := testCsiSequence{
		params:        [][]uint16{{math.MaxUint16}},
		intermediates: []byte{},
		ignore:        false,
		rune:          'm',
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[9223372036854775808m") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestCsiReset(t *testing.T) {
	expected := testCsiSequence{
		params:        [][]uint16{{1049}},
		intermediates: []byte{'?'},
		ignore:        false,
		rune:          'h',
	}
	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[3;1\x1b[?1049h") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestCsiSubparams(t *testing.T) {
	expected := testCsiSequence{
		params: [][]uint16{
			{38, 2, 255, 0, 255},
			{1},
		},
		intermediates: []byte{},
		ignore:        false,
		rune:          'm',
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[38:2:255:0:255;1m") {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}

func TestParamsBufferFilledWithSubparams(t *testing.T) {
	input := "\x1b[::::::::::::::::::::::::::::::::x\x1b"
	expected := testCsiSequence{
		params: [][]uint16{
			{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		intermediates: []byte{},
		ignore:        true,
		rune:          'x',
	}

	dispatcher := &testDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	assert.Equal(t, expected, dispatcher.dispatched[0])
}
