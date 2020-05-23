package vtparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dcsDispatcher struct {
	dispatched    bool
	intermediates []byte
	params        []int64
	ignore        bool
	r             rune
	s             []byte
}

func (p *dcsDispatcher) Print(r rune) {}

func (p *dcsDispatcher) Execute(b byte) {}

func (p *dcsDispatcher) Put(b byte) {
	p.s = append(p.s, b)
}

func (p *dcsDispatcher) Unhook() {
	p.dispatched = true
}

func (p *dcsDispatcher) Hook(params []int64, intermediates []byte, ignore bool, r rune) {
	p.intermediates = intermediates
	p.params = params
	p.ignore = ignore
	p.r = r
	p.dispatched = true
}

func (p *dcsDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {}

func (p *dcsDispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, r rune) {}

func (p *dcsDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {}

func TestDcsMaxParams(t *testing.T) {
	dispatcher := &dcsDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1bP1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;1;p\x1b") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.True(t, dispatcher.ignore)
	assert.Equal(t, maxParams, len(dispatcher.params))

	for _, param := range dispatcher.params {
		assert.Equal(t, int64(1), param)
	}
}

func TestDcsReset(t *testing.T) {
	dispatcher := &dcsDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[3;1\x1bP1$tx\x9c") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.ignore)
	assert.Equal(t, []byte([]byte{'$'}), dispatcher.intermediates)
	assert.Equal(t, []int64([]int64{1}), dispatcher.params)
}

func TestDcsParse(t *testing.T) {
	bytes := []byte{0x1b, 0x50, 0x30, 0x3b, 0x31, 0x7c, 0x31, 0x37, 0x2f, 0x61, 0x62, 0x9c}

	dispatcher := &dcsDispatcher{}
	parser := New(dispatcher)

	for _, b := range bytes {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.Equal(t, []int64([]int64{0, 1}), dispatcher.params)
	assert.Equal(t, rune('|'), dispatcher.r)
	assert.Equal(t, []byte("17/ab"), dispatcher.s)
}
