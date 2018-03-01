package canvas

type Canvas interface {
	Width() int
	Height() int
	Contains(x, y int) bool
	Get(x, y int) rune
	Set(x, y int, r rune) bool
}
