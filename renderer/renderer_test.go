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
	c.Set(canvas.Point(0, 0), 'a')
	c.Set(canvas.Point(2, 0), 'b')
	c.Set(canvas.Point(3, 1), 'c')

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
