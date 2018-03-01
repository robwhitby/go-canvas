package renderer

import (
	"io"

	"strings"

	"github.com/robwhitby/go-canvas/canvas"
)

type Renderer func(canvas.Canvas, io.Writer)

func Console(c canvas.Canvas, out io.Writer) {
	newline := []byte("\n")
	hBorder := []byte(strings.Repeat("-", c.Width()+2))
	vBorder := []byte("|")

	out.Write(hBorder)
	out.Write(newline)

	for y := 0; y < c.Height(); y++ {
		out.Write(vBorder)
		for x := 0; x < c.Width(); x++ {
			out.Write([]byte(string(c.Get(canvas.Point(x, y)))))
		}
		out.Write(vBorder)
		out.Write(newline)
	}

	out.Write(hBorder)
	out.Write(newline)
}