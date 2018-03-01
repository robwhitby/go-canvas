package renderer

import (
	"testing"

	"bytes"

	"strings"

	"github.com/robwhitby/go-canvas/canvas"
	"github.com/stretchr/testify/assert"
)

func TestConsole(t *testing.T) {
	c := canvas.NewMapCanvas(4, 2)
	c.Set(1, 1, 'a')
	c.Set(3, 1, 'b')
	c.Set(4, 2, 'c')

	out := &bytes.Buffer{}
	Console(c, out)

	expected := makeOutput(`
------
|a b |
|   c|
------
`)

	assert.Equal(t, expected, out.String())
}

func makeOutput(s string) string {
	return strings.TrimSpace(s) + "\n"
}
