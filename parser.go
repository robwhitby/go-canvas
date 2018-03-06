package go_canvas

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	lineCharacter   = 'x'
	errParseCommand = errors.New("unknown command")
)

type ParseResult struct {
	Success Command
	Failure error
}

func (result ParseResult) Apply(canvas Canvas) error {
	if result.Failure == nil {
		return result.Success.Apply(canvas)
	}
	return result.Failure
}

func ParseCommand(input string) ParseResult {
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
		return ParseResult{Success: NewCreateCanvasCommand(ints[0], ints[1])}
	}

	lineRe := regexp.MustCompile(`^[Ll] (\d+) (\d+) (\d+) (\d+)$`)
	matches = lineRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:])
		return ParseResult{Success: NewLineCommand(ints[0], ints[1], ints[2], ints[3], lineCharacter)}
	}

	rectangleRe := regexp.MustCompile(`^[Rr] (\d+) (\d+) (\d+) (\d+)$`)
	matches = rectangleRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:])
		return ParseResult{Success: NewRectangleCommand(ints[0], ints[1], ints[2], ints[3], lineCharacter)}
	}

	fillRe := regexp.MustCompile(`^[Bb] (\d+) (\d+) (\S)$`)
	matches = fillRe.FindStringSubmatch(input)
	if len(matches) > 0 {
		ints := toInt(matches[1:3])
		return ParseResult{Success: NewFillCommand(ints[0], ints[1], rune(matches[3][0]))}
	}

	if strings.ToLower(input) == "c" {
		return ParseResult{Success: NewClearCommand()}
	}

	return ParseResult{Failure: errParseCommand}
}
