package canvas

type point struct {
	x, y int
}

type MapCanvas struct {
	width  int
	height int
	points map[point]rune
}

func NewMapCanvas(width int, height int) Canvas {
	return &MapCanvas{
		width:  width,
		height: height,
		points: make(map[point]rune),
	}
}

func (c *MapCanvas) Width() int {
	return c.width
}

func (c *MapCanvas) Height() int {
	return c.height
}

func (c *MapCanvas) Contains(x, y int) bool {
	return x < c.width && y < c.height
}

func (c *MapCanvas) Get(x, y int) rune {
	val, exists := c.points[point{x, y}]
	if exists {
		return val
	}
	return ' '
}

func (c *MapCanvas) Set(x, y int, r rune) bool {
	if !c.Contains(x, y) {
		return false
	}
	c.points[point{x, y}] = r
	return true
}
