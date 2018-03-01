package main

import (
	"os"

	"github.com/robwhitby/go-canvas/canvas"
	"github.com/robwhitby/go-canvas/renderer"
)

func main() {
	c := canvas.NewMapCanvas(11, 6)
	c.Set(5, 3, '?')

	renderer.Console(c, os.Stdout)
}
