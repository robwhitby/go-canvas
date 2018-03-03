package go_canvas

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	lineCharacter   = 'x'
	errParseCommand = errors.New("unknown command")
)

func ParseCommand(input string) (Command, error) {
	toInt := func(in []string) (out []int) {
		for _, s := range in {
			i, _ := strconv.Atoi(s)
			out = append(out, i)
		}
		return
	}

	createRe := regexp.MustCompile(`^[Cc] (\d+) (\d+)$`)
	matches := createRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:])
		return NewCreateCanvasCommand(ints[0], ints[1]), nil
	}

	lineRe := regexp.MustCompile(`^[Ll] (\d+) (\d+) (\d+) (\d+)$`)
	matches = lineRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:])
		return NewLineCommand(ints[0], ints[1], ints[2], ints[3], lineCharacter), nil
	}

	rectangleRe := regexp.MustCompile(`^[Rr] (\d+) (\d+) (\d+) (\d+)$`)
	matches = rectangleRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:])
		return NewRectangleCommand(ints[0], ints[1], ints[2], ints[3], lineCharacter), nil
	}

	fillRe := regexp.MustCompile(`^[Bb] (\d+) (\d+) (\S)$`)
	matches = fillRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:3])
		return NewFillCommand(ints[0], ints[1], rune(matches[3][0])), nil
	}

	return nil, errParseCommand
}
