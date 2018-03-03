package go_canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {

	validInputs := []string{
		"C 1 1", "C 1 2", "c 40 5",
		"L 1 2 3 4", "l 1 10 100 12345",
		"R 1 2 3 4", "r 1 222 3 42",
		"B 1 2 o", "b 1 2 o",
		"C", "c",
	}

	for _, input := range validInputs {
		command, err := ParseCommand(input)
		assert.Nil(t, err, input)
		assert.NotNil(t, command, input)
	}

	invalidInputs := []string{
		"", " ", "\t", "foo", " C 1 1", "C 1 1 ",
		"C 1", "C -1 5", "C x 1", "C 1 2 3",
		"L 1 2 3", "L 1 2 3 x", "L 1 2 3 4 5",
		"R 1 2 3", "R 1 2 3 x", "R 1 3 2 4 5",
		"B 1 2", "B 1 2 3 x", "B 1 x o",
	}

	for _, input := range invalidInputs {
		command, err := ParseCommand(input)
		assert.Equal(t, errParseCommand, err, input)
		assert.Nil(t, command, input)
	}
}
