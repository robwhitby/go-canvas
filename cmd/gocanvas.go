package main

import (
	"os"

	"fmt"

	"bufio"

	"strings"

	. "github.com/robwhitby/go-canvas"
)

func main() {
	canvas := NewMapCanvas(0, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("enter command: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "q" {
			break
		}
		command, err := ParseCommand(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err := command.Apply(canvas); err != nil {
			fmt.Println(err)
			continue
		}
		StringRenderer(canvas, os.Stdout)
	}

}
