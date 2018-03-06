package main

import (
	"os"

	"fmt"

	"bufio"

	"strings"

	. "github.com/robwhitby/go-canvas"
)

func main() {
	canvas := NewCanvas(0, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("enter command: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}
		if err := ParseCommand(input).Apply(canvas); err != nil {
			fmt.Println(err)
			continue
		}
		StringRenderer(canvas, os.Stdout)
	}
}
