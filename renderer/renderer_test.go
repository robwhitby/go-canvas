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
	c.Set(0, 0, 'a')
	c.Set(2, 0, 'b')
	c.Set(3, 1, 'c')

	out := &bytes.Buffer{}
	Console(c, out)

	expected := makeExpected(`
------
|a b |
|   c|
------
`)

	assert.Equal(t, expected, out.String())
}

func makeExpected(s string) string {
	return strings.TrimSpace(s) + "\n"
}
