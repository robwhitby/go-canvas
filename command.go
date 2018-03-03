package go_canvas

import (
	"errors"
)

type Command interface {
	Apply(Canvas) error
}

type createCanvasCommand struct {
	width, height int
}

func NewCreateCanvasCommand(width, height int) createCanvasCommand {
	return createCanvasCommand{width: width, height: height}
}

func (c createCanvasCommand) Apply(canvas Canvas) error {
	canvas.Recreate(c.width, c.height)
	return nil
}

type lineCommand struct {
	x1, y1, x2, y2 int
	colour         rune
}

var (
	errLineHV = errors.New("lines must be horizontal or vertical")
)

func NewLineCommand(x1, y1, x2, y2 int, colour rune) lineCommand {
	return lineCommand{x1: x1, y1: y1, x2: x2, y2: y2, colour: colour}
}

func (l lineCommand) Apply(canvas Canvas) error {
	x1, x2 := sort(l.x1, l.x2)
	y1, y2 := sort(l.y1, l.y2)

	if !(x1 == x2 || y1 == y2) {
		return errLineHV
	}

	for y := y1; y <= y2; y++ {
		if y > canvas.Height() {
			break
		}
		for x := x1; x <= x2; x++ {
			if x > canvas.Width() {
				break
			}
			canvas.Set(x, y, l.colour)
		}
	}
	return nil
}

func sort(i, j int) (int, int) {
	if i > j {
		return j, i
	}
	return i, j
}

type rectangleCommand struct {
	x1, y1, x2, y2 int
	colour         rune
}

func NewRectangleCommand(x1, y1, x2, y2 int, colour rune) rectangleCommand {
	return rectangleCommand{x1: x1, y1: y1, x2: x2, y2: y2, colour: colour}
}

func (c rectangleCommand) Apply(canvas Canvas) error {
	NewLineCommand(c.x1, c.y1, c.x2, c.y1, c.colour).Apply(canvas) //top
	NewLineCommand(c.x1, c.y2, c.x2, c.y2, c.colour).Apply(canvas) //bottom
	NewLineCommand(c.x1, c.y1, c.x1, c.y2, c.colour).Apply(canvas) //left
	NewLineCommand(c.x2, c.y1, c.x2, c.y2, c.colour).Apply(canvas) //right
	return nil
}

type fillCommand struct {
	x, y   int
	colour rune
}

func NewFillCommand(x, y int, colour rune) fillCommand {
	return fillCommand{x: x, y: y, colour: colour}
}

func (f fillCommand) Apply(canvas Canvas) error {
	startColour := canvas.Get(f.x, f.y)
	if startColour == f.colour {
		return nil
	}
	pointsSeen := map[point]struct{}{
		{f.x, f.y}: {},
	}
	q := []point{{f.x, f.y}}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		canvas.Set(p.x, p.y, f.colour)

		neighbours := []point{
			{p.x + 1, p.y},
			{p.x - 1, p.y},
			{p.x, p.y + 1},
			{p.x, p.y - 1},
		}
		for _, n := range neighbours {
			_, seen := pointsSeen[n]
			if !seen && canvas.Contains(n.x, n.y) && canvas.Get(n.x, n.y) == startColour {
				q = append(q, n)
				pointsSeen[n] = struct{}{}
			}
		}
	}
	return nil
}

type clearCommand struct{}

func NewClearCommand() clearCommand {
	return clearCommand{}
}

func (clearCommand) Apply(canvas Canvas) error {
	canvas.Recreate(canvas.Width(), canvas.Height())
	return nil
}
