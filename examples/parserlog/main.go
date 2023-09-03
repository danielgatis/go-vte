package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/danielgatis/go-vte"
)

type dispatcher struct{}

func (p *dispatcher) Print(r rune) {
	fmt.Printf("[Print] %c\n", r)
}

func (p *dispatcher) Execute(b byte) {
	fmt.Printf("[Execute] %02x\n", b)
}

func (p *dispatcher) Put(b byte) {
	fmt.Printf("[Put] %02x\n", b)
}

func (p *dispatcher) Unhook() {
	fmt.Printf("[Unhook]\n")
}

func (p *dispatcher) Hook(params *vte.Params, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[Hook] params=%v, intermediates=%v, ignore=%v, r=%c\n", params, intermediates, ignore, r)
}

func (p *dispatcher) OscDispatch(params [][]byte, bellTerminated bool) {
	fmt.Printf("[OscDispatch] params=%v, bellTerminated=%v\n", params, bellTerminated)
}

func (p *dispatcher) CsiDispatch(params *vte.Params, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[CsiDispatch] params=%v, intermediates=%v, ignore=%v, r=%c\n", params, intermediates, ignore, r)
}

func (p *dispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {
	fmt.Printf("[EscDispatch] intermediates=%v, ignore=%v, byte=%02x\n", intermediates, ignore, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	dispatcher := &dispatcher{}
	parser := vte.New(dispatcher)

	buff := make([]byte, 2048)

	for {
		n, err := reader.Read(buff)

		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Printf("Err %v:", err)
			return
		}

		for _, b := range buff[:n] {
			parser.Advance(b)
		}
	}
}
