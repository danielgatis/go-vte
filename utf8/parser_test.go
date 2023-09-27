package utf8

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	output := ""

	parser := New(
		func(r rune) {
			output = output + string(r)
		},

		func() {
			output = output + "�"
		},
	)

	data, err := os.ReadFile("../fixtures/UTF-8-demo.txt")

	if err != nil {
		t.Fatal(err)
	}

	for _, b := range data {
		parser.Advance(b)
	}

	assert.Equal(t, string(data), output)
}
