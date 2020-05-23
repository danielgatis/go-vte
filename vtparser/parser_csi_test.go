package vtparser

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type csiDispatcher struct {
	dispatched    bool
	intermediates []byte
	params        []int64
	ignore        bool
}

func (p *csiDispatcher) Print(r rune) {}

func (p *csiDispatcher) Execute(b byte) {}

func (p *csiDispatcher) Put(b byte) {}

func (p *csiDispatcher) Unhook() {}

func (p *csiDispatcher) Hook(params []int64, intermediates []byte, ignore bool, r rune) {}

func (p *csiDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {}

func (p *csiDispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, r rune) {
	p.intermediates = intermediates
	p.params = params
	p.ignore = ignore
	p.dispatched = true
}

func (p *csiDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {}

func TestCsiMaxParams(t *testing.T) {
	strParams := "\x1b["

	for i := 0; i < maxParams-1; i++ {
		strParams += "1;"
	}

	strParams += "p"

	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(strParams) {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.ignore)
	assert.Equal(t, maxParams, len(dispatcher.params))
}

func TestCsiParamsIgnoreLong(t *testing.T) {
	strParams := "\x1b["

	for i := 0; i < maxParams; i++ {
		strParams += "1;"
	}

	strParams += "p"

	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte(strParams) {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.True(t, dispatcher.ignore)
	assert.Equal(t, maxParams, len(dispatcher.params))
}

func TestCsiParamsTrailingSemicolon(t *testing.T) {
	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[4;m") {
		parser.Advance(b)
	}

	assert.Equal(t, []int64{4, 0}, dispatcher.params)
}

func TestCsiSemiSetUnderline(t *testing.T) {
	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[;4m") {
		parser.Advance(b)
	}

	assert.Equal(t, []int64{0, 4}, dispatcher.params)
}

func TestLongCsiParam(t *testing.T) {
	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[9223372036854775808m") {
		parser.Advance(b)
	}

	assert.Equal(t, []int64{math.MaxInt64}, dispatcher.params)
}

func TestLongCsiReset(t *testing.T) {
	dispatcher := &csiDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[3;1\x1b[?1049h") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.ignore)
	assert.Equal(t, []byte{'?'}, dispatcher.intermediates)
	assert.Equal(t, []int64{1049}, dispatcher.params)
}
