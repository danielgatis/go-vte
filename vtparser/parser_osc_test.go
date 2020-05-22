package vtparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type oscDispatcher struct {
	dispatched     bool
	bellTerminated bool
	params         [][]byte
}

func (p *oscDispatcher) Print(char uint32) {}

func (p *oscDispatcher) Execute(b byte) {}

func (p *oscDispatcher) Put(b byte) {}

func (p *oscDispatcher) Unhook() {}

func (p *oscDispatcher) Hook(params []int64, intermediates []byte, ignore bool, char uint32) {}

func (p *oscDispatcher) OscDispatch(params [][]byte, bellTerminated bool) {
	p.dispatched = true
	p.bellTerminated = bellTerminated
	p.params = params
}

func (p *oscDispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, char uint32) {}

func (p *oscDispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {}

func TestOsc(t *testing.T) {
	var oscBytes = []byte{
		0x1b, 0x5d, // Begin OSC
		'2', ';', 'j', 'w', 'i', 'l', 'm', '@', 'j', 'w', 'i', 'l', 'm', '-', 'd',
		'e', 's', 'k', ':', ' ', '~', '/', 'c', 'o', 'd', 'e', '/', 'a', 'l', 'a',
		'c', 'r', 'i', 't', 't', 'y', 0x07, // End OSC
	}

	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range oscBytes {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.True(t, dispatcher.bellTerminated)
	assert.Equal(t, 2, len(dispatcher.params))
	assert.Equal(t, oscBytes[2:3], dispatcher.params[0])
	assert.Equal(t, oscBytes[4:len(oscBytes)-1], dispatcher.params[1])
}

func TestEmptyOsc(t *testing.T) {
	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte{0x1b, 0x5d, 0x07} {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.True(t, dispatcher.bellTerminated)
	assert.Equal(t, 1, len(dispatcher.params))
	assert.Equal(t, []byte{}, dispatcher.params[0])
}

func TestOscMaxParams(t *testing.T) {
	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b];;;;;;;;;;;;;;;;;\x1b") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.bellTerminated)
	assert.Equal(t, maxParams, len(dispatcher.params))
	assert.Equal(t, []byte{}, dispatcher.params[0])

	for _, param := range dispatcher.params {
		assert.Equal(t, 0, len(param))
	}
}

func TestOscBellTerminated(t *testing.T) {
	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b]11;ff/00/ff\x07") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.True(t, dispatcher.bellTerminated)
}

func TestOscC0StTerminated(t *testing.T) {
	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range []byte("\x1b]11;ff/00/ff\x1b\\") {
		parser.Advance(b)
	}

	assert.True(t, dispatcher.dispatched)
	assert.False(t, dispatcher.bellTerminated)
}

func TestOscWithUTF8Arguments(t *testing.T) {
	bytes := []byte{
		0x0d, 0x1b, 0x5d, 0x32, 0x3b, 0x65, 0x63, 0x68, 0x6f, 0x20, 0x27, 0xc2, 0xaf, 0x5c,
		0x5f, 0x28, 0xe3, 0x83, 0x84, 0x29, 0x5f, 0x2f, 0xc2, 0xaf, 0x27, 0x20, 0x26, 0x26,
		0x20, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x20, 0x31, 0x07,
	}

	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range bytes {
		parser.Advance(b)
	}

	assert.Equal(t, []uint8([]byte{'2'}), dispatcher.params[0])
	assert.Equal(t, bytes[5:(len(bytes)-1)], dispatcher.params[1])
}

func TestOscContainingStringTerminator(t *testing.T) {
	bytes := []byte("\x1b]2;\xe6\x9c\xab\x1b\\")

	dispatcher := &oscDispatcher{}
	parser := New(dispatcher)

	for _, b := range bytes {
		parser.Advance(b)
	}

	assert.Equal(t, bytes[4:(len(bytes)-2)], dispatcher.params[1])
}

func TestOcsExceedMaxBufferSize(t *testing.T) {
	numBytes := maxOscRaw + 100
	inputStart := []byte{0x1b, ']', '5', '2', ';', 's'}
	inputEnd := []byte{0x07}

	dispatcher := &oscDispatcher{}
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

	assert.True(t, dispatcher.dispatched)
	assert.Equal(t, 2, len(dispatcher.params))
	assert.Equal(t, []byte("52"), dispatcher.params[0])
	assert.Equal(t, numBytes+len(inputEnd), len(dispatcher.params[1]))
}
