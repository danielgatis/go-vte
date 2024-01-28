package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/danielgatis/go-vte"
)

var _ (vte.Performer) = (*performer)(nil)

type performer struct{}

func (p *performer) Print(r rune) {
	fmt.Printf("[Print] %c\n", r)
}

func (p *performer) Execute(b byte) {
	fmt.Printf("[Execute] %02x\n", b)
}

func (p *performer) Put(b byte) {
	fmt.Printf("[Put] %02x\n", b)
}

func (p *performer) Unhook() {
	fmt.Printf("[Unhook]\n")
}

func (p *performer) Hook(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[Hook] params=%v, intermediates=%v, ignore=%v, r=%c\n", params, intermediates, ignore, r)
}

func (p *performer) OscDispatch(params [][]byte, bellTerminated bool) {
	fmt.Printf("[OscDispatch] params=%v, bellTerminated=%v\n", params, bellTerminated)
}

func (p *performer) CsiDispatch(params [][]uint16, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[CsiDispatch] params=%v, intermediates=%v, ignore=%v, r=%c\n", params, intermediates, ignore, r)
}

func (p *performer) EscDispatch(intermediates []byte, ignore bool, b byte) {
	fmt.Printf("[EscDispatch] intermediates=%v, ignore=%v, byte=%02x\n", intermediates, ignore, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	performer := &performer{}
	parser := vte.NewParser(performer)

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
