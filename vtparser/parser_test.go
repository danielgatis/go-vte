package vtparser

import (
	"io/ioutil"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type benchDispatcher struct{}

func (p *benchDispatcher) Print(r rune) {}

func (p *benchDispatcher) Execute(b byte) {}

func (p *benchDispatcher) Put(b byte) {}

func (p *benchDispatcher) Unhook() {}

func (p *benchDispatcher) Hook(params []int64, intermediates []byte, ignore bool, r rune) {}

func (p *benchDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {}

func (p *benchDispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, r rune) {
}

func (p *benchDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {}

func BenchmarkNext(bm *testing.B) {
	bytes, err := ioutil.ReadFile("../fixtures/demo.vte")

	if err != nil {
		bm.Fatalf("Error: %v", err)
	}

	bm.ResetTimer()
	dispatcher := &benchDispatcher{}
	parser := New(dispatcher)

	for _, b := range bytes {
		parser.Advance(b)
	}
}

func TestSadd64(t *testing.T) {
	assert.Equal(t, int64(math.MaxInt64), sadd64(math.MaxInt64, 1))
	assert.Equal(t, int64(math.MinInt64), sadd64(math.MinInt64, -1))
}

func TestSmul64(t *testing.T) {
	assert.Equal(t, int64(math.MaxInt64), smul64(math.MaxInt64, 2))
	assert.Equal(t, int64(math.MinInt64), smul64(math.MaxInt64, -2))
	assert.Equal(t, int64(math.MinInt64), smul64(math.MinInt64, 2))
	assert.Equal(t, int64(math.MaxInt64), smul64(math.MinInt64, -2))
}
