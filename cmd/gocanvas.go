package main

import (
	"os"

	"github.com/robwhitby/go-canvas/canvas"
	"github.com/robwhitby/go-canvas/renderer"
)

func main() {
	c := canvas.NewMapCanvas(5, 3)
	c.Set(3, 2, '?')

	renderer.Console(c, os.Stdout)
}
