# go-vte

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-vte?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-vte)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-vte/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-vte)

A GO version of https://github.com/alacritty/vte.

The pkg `vtparse` implements a state machine that mirrors the behaviour of DEC (Digital Equipment Corporation) VT hardware terminals. The state machine was originally described by Paul Williams; more information can be found here: http://www.vt100.net/emu/dec_ansi_parser.

## Install

```bash
go get -u github.com/danielgatis/go-vte
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-vte"
```

### Example

Please look at: [examples/parserlog/main.go](examples/parserlog/main.go)

```
❯ echo -ne "Hello\nWorld" | go run ./examples/parserlog/main.go
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
