package go_canvas

import (
	"io"

	"strings"
)

type Renderer func(Canvas, io.Writer)

func StringRenderer(canvas Canvas, out io.Writer) {
	if canvas.Width() == 0 || canvas.Height() == 0 {
		return
	}
	newline := []byte("\n")
	hBorder := []byte(strings.Repeat("-", canvas.Width()+2))
	vBorder := []byte("|")

	out.Write(hBorder)
	out.Write(newline)

	for y := 1; y <= canvas.Height(); y++ {
		out.Write(vBorder)
		for x := 1; x <= canvas.Width(); x++ {
			out.Write([]byte(string(canvas.Get(x, y))))
		}
		out.Write(vBorder)
		out.Write(newline)
	}

	out.Write(hBorder)
	out.Write(newline)
}
