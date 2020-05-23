# go-vte

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-vte?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-vte)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-vte/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-vte)
[![Release](https://img.shields.io/github/release/danielgatis/go-vte.svg?style=flat-square)](https://github.com/danielgatis/go-vte/releases/latest)

A GO version of https://github.com/alacritty/vte.

The pkg `vtparse` implements a state machine that mirrors the behaviour of DEC (Digital Equipment Corporation) VT hardware terminals. The state machine was originally described by Paul Williams; more information can be found here: http://www.vt100.net/emu/dec_ansi_parser.

The pkg `utf8` implements a state machine that reads UTF-8.

## Install

```bash
go get -u github.com/danielgatis/go-vte
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-vte/vtparser"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/danielgatis/go-vte/vtparser"
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

func (p *dispatcher) Hook(params []int64, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[Hook] params=%v, intermediates=%v, ignore=%v, r=%v\n", params, intermediates, ignore, r)
}

func (p *dispatcher) OscDispatch(params [][]byte, bellTerminated bool) {
	fmt.Printf("[OscDispatch] params=%v, bellTerminated=%v\n", params, bellTerminated)
}

func (p *dispatcher) CsiDispatch(params []int64, intermediates []byte, ignore bool, r rune) {
	fmt.Printf("[CsiDispatch] params=%v, intermediates=%v, ignore=%v, r=%v\n", params, intermediates, ignore, r)
}

func (p *dispatcher) EscDispatch(intermediates []byte, ignore bool, b byte) {
	fmt.Printf("[EscDispatch] intermediates=%v, ignore=%v, byte=%02x\n", intermediates, ignore, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	parser := vtparser.New(&dispatcher{})
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

```


```
❯ echo -ne "Hello\nWorld" | go run main.go
[Print] H
[Print] e
[Print] l
[Print] l
[Print] o
[Execute] 0a
[Print] W
[Print] o
[Print] r
[Print] l
[Print] d
```


## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)
