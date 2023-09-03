package vte

import (
	"fmt"
	"strings"
)

const maxParams = 32

// Params is a parameters for VTE.
type Params struct {
	// Number of subparameters in each parameter
	//
	// For each entry in the `params` slice, this stores the length of the param as number of
	// subparams at the same index as the param in the `params` slice.
	//
	// At the subparam positions the length will always be `0`.
	subparams [maxParams]uint8

	// All parameters and subparameters
	params [maxParams]uint16

	// Number of subparams in the current parameter.
	current_subparams uint8

	len int
}

// NewParams creates a new Params.
func NewParams() *Params {
	p := new(Params)
	return p
}

// Len returns the number of parameters.
func (p *Params) Len() int {
	return p.len
}

// IsEmpty returns true if there are no parameters.
func (p *Params) IsEmpty() bool {
	return p.len == 0
}

// IsFull returns true if there are no more parameters can be added.
func (p *Params) IsFull() bool {
	return p.len == maxParams
}

// Clear clears all parameters.
func (p *Params) Clear() {
	p.current_subparams = 0
	p.len = 0
}

// Push pushes a parameter.
func (p *Params) Push(param uint16) {
	p.subparams[p.len-int(p.current_subparams)] = p.current_subparams + 1
	p.params[p.len] = param
	p.current_subparams = 0
	p.len++
}

// Extend extends the last parameter.
func (p *Params) Extend(param uint16) {
	p.subparams[p.len-int(p.current_subparams)] = p.current_subparams + 1
	p.params[p.len] = param
	p.current_subparams++
	p.len++
}

// Range iterates over all parameters.
func (p *Params) Range(f func(param []uint16)) {
	for i := 0; i < p.len; {
		numSubparams := p.subparams[i]
		param := p.params[i : i+int(numSubparams)]
		i += int(numSubparams)
		f(param)
	}
}

// Params returns all parameters and their subparameters as a slice.
func (p *Params) Params() []uint16 {
	var params []uint16
	p.Range(func(param []uint16) {
		params = append(params, param...)
	})
	return params
}

// String returns a string representation of the parameters.
func (p *Params) String() string {
	var b strings.Builder
	b.WriteString("[")
	i := 0
	p.Range(func(param []uint16) {
		if i > 0 {
			b.WriteRune(';')
		}
		for j, subparam := range param {
			if j > 0 {
				b.WriteRune(':')
			}
			b.WriteString(fmt.Sprint(subparam))
		}
		i++
	})
	b.WriteString("]")
	return b.String()
}
