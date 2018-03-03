package main

import (
	"os"

	"fmt"

	"bufio"

	. "github.com/robwhitby/go-canvas"
)

const prompt = "enter command: "

func main() {
	canvas := NewMapCanvas(0, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()

		if input == "Q" {
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
