package canvas

import "errors"

type point struct {
	x, y int
}

func Point(x int, y int) point {
	return point{x: x, y: y}
}

type Canvas interface {
	Width() int
	Height() int
	Contains(point) bool
	Get(point) rune
	Set(point, rune) bool
}

var (
	ErrOOB = errors.New("out of bounds")
)
