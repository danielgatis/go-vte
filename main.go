package main

//go:generate ruby ./utf8/_tablegen.rb
//go:generate ruby ./vtparser/_tablegen.rb

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/danielgatis/go-vte/vtparser"
)

type performer struct{}

func (p *performer) Print(char uint32) {
	fmt.Printf("[Print] %c\n", char)
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

func (p *performer) Hook(params []int64, intermediates []byte, ignore bool, char uint32) {
	fmt.Printf("[Hook] params=%v, intermediates=%v, ignore=%v, char=%c\n", params, intermediates, ignore, char)
}

func (p *performer) OscDispatch(params [][]byte, bellTerminated bool) {
	fmt.Printf("[OscDispatch] params=%v, bellTerminated=%v\n", params, bellTerminated)
}

func (p *performer) CsiDispatch(params []int64, intermediates []byte, ignore bool, char uint32) {
	fmt.Printf("[CsiDispatch] params=%v, intermediates=%v, ignore=%v, char=%v\n", params, intermediates, ignore, char)
}

func (p *performer) EscDispatch(intermediates []byte, ignore bool, b byte) {
	fmt.Printf("[EscDispatch] intermediates=%v, ignore=%v, byte=%02x\n", intermediates, ignore, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	parser := vtparser.New(&performer{})
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
