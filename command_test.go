package go_canvas

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestNewCanvasCommand(t *testing.T) {

	t.Run("creates new canvas", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		createCanvas := NewCreateCanvasCommand(4, 5)
		err := createCanvas.Apply(canvas)

		assert.Nil(t, err)
		assert.Equal(t, 4, canvas.Width())
		assert.Equal(t, 5, canvas.Height())
	})

}

func TestLineCommand(t *testing.T) {

	t.Run("draws a horizontal line", func(t *testing.T) {
		canvas := NewCanvas(10, 1)
		line := NewLineCommand(5, 1, 7, 1, 'h')
		err := line.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "    hhh   \n")
	})

	t.Run("draws a vertical line", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		line := NewLineCommand(2, 1, 2, 3, 'v')
		err := line.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, " v \n v \n v \n")
	})

	t.Run("error when not horizontal or vertical", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		line := NewLineCommand(1, 2, 3, 4, 'd')
		err := line.Apply(canvas)

		assert.Equal(t, errLineHV, err)
		assertCanvas(t, canvas, "   \n   \n   \n")
	})

	t.Run("ignores out of bounds", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		line := NewLineCommand(1, 1, 5, 1, 'h')
		err := line.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "hhh\n   \n   \n")
	})

	t.Run("draws lines 'backwards'", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		lineH := NewLineCommand(3, 1, 1, 1, 'h')
		lineV := NewLineCommand(1, 3, 1, 1, 'v')

		err := lineH.Apply(canvas)
		assert.Nil(t, err)
		err = lineV.Apply(canvas)
		assert.Nil(t, err)
		assertCanvas(t, canvas, "vhh\nv  \nv  \n")
	})
}

func TestRectangleCommand(t *testing.T) {

	t.Run("draws a rectangle", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		rectangle := NewRectangleCommand(3, 3, 1, 1, 'r')
		err := rectangle.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "rrr\nr r\nrrr\n")
	})
}

func TestFillCommand(t *testing.T) {

	t.Run("fills an empty canvas", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		fill := NewFillCommand(1, 1, 'f')
		err := fill.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "fff\nfff\nfff\n")
	})

	t.Run("fills an area", func(t *testing.T) {
		canvas := NewCanvas(5, 5)
		rectangle := NewRectangleCommand(2, 2, 4, 4, 'r')
		rectangle.Apply(canvas)

		fill := NewFillCommand(1, 1, 'f')
		err := fill.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "fffff\nfrrrf\nfr rf\nfrrrf\nfffff\n")

		fill = NewFillCommand(2, 3, '.')
		err = fill.Apply(canvas)

		assert.Nil(t, err)
		assertCanvas(t, canvas, "fffff\nf...f\nf. .f\nf...f\nfffff\n")
	})
}

func TestClearCommand(t *testing.T) {
	t.Run("clears the canvas", func(t *testing.T) {
		canvas := NewCanvas(3, 3)
		canvas.Set(2, 3, 's')

		NewClearCommand().Apply(canvas)
		assertCanvas(t, canvas, "   \n   \n   \n")
	})
}

func assertCanvas(t *testing.T, canvas Canvas, expected string) {
	t.Helper()
	var actual bytes.Buffer
	for y := 1; y <= canvas.Height(); y++ {
		for x := 1; x <= canvas.Width(); x++ {
			actual.Write([]byte(string(canvas.Get(x, y))))
		}
		actual.Write([]byte("\n"))
	}
	assert.Equal(t, expected, actual.String())
}
