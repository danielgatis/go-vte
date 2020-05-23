package vtparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ecsDispatcher struct {
	dispatched    bool
	intermediates []byte
	b             byte
	ignore        bool
}

func (p *ecsDispatcher) Print(r rune) {}

func (p *ecsDispatcher) Execute(b byte) {}

func (p *ecsDispatcher) Put(b byte) {}

func (p *ecsDispatcher) Unhook() {}

func (p *ecsDispatcher) Hook(params []int64, intermediates []byte, ignore bool, r rune) {}

func (p *ecsDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {}

func (p *ecsDispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, r rune) {}

func (p *ecsDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {
	p.intermediates = intermediates
	p.b = b
	p.ignore = ignore
	p.dispatched = true
}

func TestEscReset(t *testing.T) {
	dispatcher := &ecsDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b[3;1\x1b(A") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.ignore)
	assert.Equal(t, []byte([]byte{'('}), dispatcher.intermediates)
}
