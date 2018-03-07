package go_canvas

type Canvas interface {
	Width() int
	Height() int
	Contains(x, y int) bool
	Get(x, y int) rune
	Set(x, y int, r rune)
	Recreate(width, height int)
}

func NewCanvas(width int, height int) Canvas {
	//canvas := &MapCanvas{}
	canvas := &ArrayCanvas{}
	canvas.Recreate(width, height)
	return canvas
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

// array canvas

type ArrayCanvas struct {
	width  int
	height int
	points [][]rune
}

func (c *ArrayCanvas) Width() int {
	return c.width
}

func (c *ArrayCanvas) Height() int {
	return c.height
}

func (c *ArrayCanvas) Contains(x, y int) bool {
	return x > 0 && x <= c.width && y > 0 && y <= c.height
}

func (c *ArrayCanvas) Get(x, y int) rune {
	return c.points[y-1][x-1]
}

func (c *ArrayCanvas) Set(x, y int, r rune) {
	c.points[y-1][x-1] = r
}

func (c *ArrayCanvas) Recreate(width, height int) {
	c.width = width
	c.height = height
	c.points = make([][]rune, height)
	for y := range c.points {
		c.points[y] = make([]rune, width)
		for x := range c.points[y] {
			c.points[y][x] = ' '
		}
	}
}
