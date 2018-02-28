package main

import (
	"fmt"
	"github.com/robwhitby/go-canvas/canvas"
)


func main() {
	c := canvas.NewMapCanvas(20,10)
	fmt.Println(c)
}
