package vte

import (
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testSequence interface {
	sequence()
}

type testCsiSequence struct {
	params        [][]uint16
	intermediates []byte
	ignore        bool
	rune          rune
}

func (testCsiSequence) sequence() {}

type testOscSequence struct {
	params [][]byte
	bell   bool
}

func (testOscSequence) sequence() {}

type testEscSequence struct {
	intermediates []byte
	ignore        bool
	b             byte
}

func (testEscSequence) sequence() {}

type testDcsHookSequence struct {
	params        [][]uint16
	intermediates []byte
	ignore        bool
	rune          rune
}

func (testDcsHookSequence) sequence() {}

type testDcsPutSequence byte

func (testDcsPutSequence) sequence() {}

type testDcsUnhookSequence struct{}

func (testDcsUnhookSequence) sequence() {}

type testDispatcher struct {
	dispatched []testSequence
}

var _ Performer = &testDispatcher{}

// CsiDispatch implements Performer.
func (d *testDispatcher) CsiDispatch(params *Params, intermediates []byte, ignore bool, r rune) {
	var csiParams [][]uint16
	params.Range(func(param []uint16) {
		csiParams = append(csiParams, param)
	})
	d.dispatched = append(d.dispatched, testCsiSequence{
		params:        csiParams,
		intermediates: intermediates,
		ignore:        ignore,
		rune:          r,
	})
}

// EscDispatch implements Performer.
func (d *testDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {
	d.dispatched = append(d.dispatched, testEscSequence{
		intermediates: intermediates,
		ignore:        ignore,
		b:             b,
	})
}

// Execute implements Performer.
func (*testDispatcher) Execute(b byte) {}

// Hook implements Performer.
func (d *testDispatcher) Hook(params *Params, intermediates []byte, ignore bool, r rune) {
	var hookParams [][]uint16
	params.Range(func(param []uint16) {
		hookParams = append(hookParams, param)
	})
	d.dispatched = append(d.dispatched, testDcsHookSequence{
		params:        hookParams,
		intermediates: intermediates,
		ignore:        ignore,
		rune:          r,
	})
}

// OscDispatch implements Performer.
func (d *testDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {
	d.dispatched = append(d.dispatched, testOscSequence{
		params: params,
		bell:   bellTerminated,
	})
}

// Print implements Performer.
func (*testDispatcher) Print(r rune) {}

// Put implements Performer.
func (d *testDispatcher) Put(b byte) {
	d.dispatched = append(d.dispatched, testDcsPutSequence(b))
}

// Unhook implements Performer.
func (d *testDispatcher) Unhook() {
	d.dispatched = append(d.dispatched, testDcsUnhookSequence{})
}

func TestSaddu16(t *testing.T) {
	assert.Equal(t, uint16(math.MaxUint16), saddu16(math.MaxUint16, 1))
	assert.Equal(t, uint16(1), saddu16(1, 0))
}

func TestSmulu16(t *testing.T) {
	assert.Equal(t, uint16(math.MaxUint16), smulu16(math.MaxUint16, 1))
	assert.Equal(t, uint16(math.MaxUint16), smulu16(math.MaxUint16, 2))
	assert.Equal(t, uint16(0), smulu16(math.MaxUint16, 0))
	assert.Equal(t, uint16(0), smulu16(0, 0))
}

type benchDispatcher struct{}

func (p *benchDispatcher) Print(r rune) {}

func (p *benchDispatcher) Execute(b byte) {}

func (p *benchDispatcher) Put(b byte) {}

func (p *benchDispatcher) Unhook() {}

func (p *benchDispatcher) Hook(params *Params, intermediates []byte, ignore bool, r rune) {}

func (p *benchDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {}

func (p *benchDispatcher) CsiDispatch(params *Params, intermediates []byte, ignore bool, r rune) {
}

func (p *benchDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {}

func BenchmarkNext(bm *testing.B) {
	bytes, err := os.ReadFile("./fixtures/demo.vte")

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

func BenchmarkStateChanges(bm *testing.B) {
	input := "\x1b]2;X\x1b\\ \x1b[0m \x1bP0@\x1b\\"

	for i := 0; i < bm.N; i++ {
		dispatcher := &benchDispatcher{}
		parser := New(dispatcher)

		for i := 0; i < 1000; i++ {
			for _, b := range []byte(input) {
				parser.Advance(b)
			}
		}
	}
}
