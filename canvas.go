package go_canvas

type Canvas interface {
	Width() int
	Height() int
	Contains(x, y int) bool
	Get(x, y int) rune
	Set(x, y int, r rune)
	Recreate(x, y int)
}

func NewCanvas(width int, height int) Canvas {
	return &MapCanvas{
		width:  width,
		height: height,
		points: make(map[point]rune),
	}
}

type point struct {
	x, y int
}

type MapCanvas struct {
	width  int
	height int
	points map[point]rune
}

func (c *MapCanvas) Width() int {
	return c.width
}

func (c *MapCanvas) Height() int {
	return c.height
}

func (c *MapCanvas) Contains(x, y int) bool {
	return x > 0 && x <= c.width && y > 0 && y <= c.height
}

func (c *MapCanvas) Get(x, y int) rune {
	val, exists := c.points[point{x, y}]
	if exists {
		return val
	}
	return ' '
}

func (c *MapCanvas) Set(x, y int, r rune) {
	c.points[point{x, y}] = r
}

func (c *MapCanvas) Recreate(width, height int) {
	c.width = width
	c.height = height
	c.points = make(map[point]rune)
}
