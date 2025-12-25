package vte

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for Bug 1: OSC buffer should respect maxOscRaw limit
func TestOscBufferLimit(t *testing.T) {
	// Send more bytes than maxOscRaw allows
	numBytes := maxOscRaw + 500

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	// Start OSC sequence: ESC ]
	parser.Advance(0x1b)
	parser.Advance(']')

	// Send parameter "0;"
	parser.Advance('0')
	parser.Advance(';')

	// Send numBytes of data
	for i := 0; i < numBytes; i++ {
		parser.Advance('a')
	}

	// End with BEL
	parser.Advance(0x07)

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testOscSequence)

	// The second param (data) should be limited to maxOscRaw
	// Total oscRaw = "0" + data, but param[1] is just the data part
	assert.LessOrEqual(t, len(seq.params[1]), maxOscRaw,
		"OSC buffer should be limited to maxOscRaw bytes")
}

// Test for Bug 3: HookAction should respect hasParam
func TestDcsWithoutParams(t *testing.T) {
	// DCS sequence without any parameters: ESC P q ... ESC \
	// Should result in empty params, not [[0]]
	input := "\x1bPq"

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testDcsHookSequence)

	// Without explicit parameters, params should be empty
	assert.Equal(t, 0, len(seq.params),
		"DCS without params should have empty params slice, not [[0]]")
}

// Test for Bug 3: HookAction with explicit zero param
func TestDcsWithExplicitZeroParam(t *testing.T) {
	// DCS sequence with explicit 0: ESC P 0 q
	input := "\x1bP0q"

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testDcsHookSequence)

	// With explicit 0, params should be [[0]]
	assert.Equal(t, 1, len(seq.params),
		"DCS with explicit 0 param should have params")
	assert.Equal(t, []uint16{0}, seq.params[0])
}

// Test for Bug 3: HookAction with trailing semicolon
func TestDcsWithTrailingSemicolon(t *testing.T) {
	// DCS sequence with trailing semicolon: ESC P 1; q
	// The semicolon indicates a second param exists (implicit 0)
	input := "\x1bP1;q"

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testDcsHookSequence)

	// With "1;", should have [[1], [0]] because semicolon implies second param
	assert.Equal(t, 2, len(seq.params),
		"DCS with trailing semicolon should have 2 params")
	assert.Equal(t, []uint16{1}, seq.params[0])
	assert.Equal(t, []uint16{0}, seq.params[1])
}

// Verify CSI without params has empty params (for comparison)
func TestCsiWithoutParams(t *testing.T) {
	// CSI sequence without parameters: ESC [ m
	input := "\x1b[m"

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	for _, b := range []byte(input) {
		parser.Advance(b)
	}

	assert.Equal(t, 1, len(dispatcher.dispatched))
	seq := dispatcher.dispatched[0].(testCsiSequence)

	// Without explicit parameters, params should be empty
	assert.Equal(t, 0, len(seq.params),
		"CSI without params should have empty params slice")
}

// Test for params.Push/Extend bounds safety
func TestParamsPushBoundsSafety(t *testing.T) {
	p := new(params)

	// Fill params to max
	for i := 0; i < maxParams; i++ {
		p.Push(uint16(i))
	}

	assert.True(t, p.IsFull())
	assert.Equal(t, maxParams, p.Len())

	// This should NOT panic - Push should be a no-op when full
	assert.NotPanics(t, func() {
		p.Push(999)
	})

	// Length should still be maxParams
	assert.Equal(t, maxParams, p.Len())
}

// Test for params.Extend bounds safety
func TestParamsExtendBoundsSafety(t *testing.T) {
	p := new(params)

	// Fill params to max using Extend
	for i := 0; i < maxParams; i++ {
		p.Extend(uint16(i))
	}

	assert.True(t, p.IsFull())
	assert.Equal(t, maxParams, p.Len())

	// This should NOT panic - Extend should be a no-op when full
	assert.NotPanics(t, func() {
		p.Extend(999)
	})

	// Length should still be maxParams
	assert.Equal(t, maxParams, p.Len())
}

// Test CSI with max params doesn't panic
func TestCsiMaxParamsNoPanic(t *testing.T) {
	// Build a CSI sequence with more than maxParams parameters
	// ESC [ 1;2;3;...;35 m (35 params, more than maxParams=32)
	var params []string
	for i := 1; i <= 35; i++ {
		params = append(params, "1")
	}
	input := "\x1b[" + strings.Join(params, ";") + "m"

	dispatcher := &testDispatcher{}
	parser := NewParser(dispatcher)

	// This should NOT panic
	assert.NotPanics(t, func() {
		for _, b := range []byte(input) {
			parser.Advance(b)
		}
	})

	assert.Equal(t, 1, len(dispatcher.dispatched))
}
