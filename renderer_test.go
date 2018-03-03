package go_canvas

import (
	"testing"

	"bytes"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestStringRenderer(t *testing.T) {
	canvas := NewMapCanvas(4, 2)
	canvas.Set(1, 1, 'a')
	canvas.Set(3, 1, 'b')
	canvas.Set(4, 2, 'c')

	out := &bytes.Buffer{}
	StringRenderer(canvas, out)

	expected := makeOutput(`
------
|a b |
|   c|
------
`)

	assert.Equal(t, expected, out.String())
}

func TestStringRenderer_ZeroCanvas(t *testing.T) {
	canvas := NewMapCanvas(0, 0)
	out := &bytes.Buffer{}
	StringRenderer(canvas, out)
	assert.Equal(t, "", out.String())
}

func makeOutput(s string) string {
	return strings.TrimSpace(s) + "\n"
}
