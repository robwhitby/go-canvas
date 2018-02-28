package canvas

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

func (c *MapCanvas) Contains(p point) bool {
	return p.x < c.width && p.y < c.height
}

func (c *MapCanvas) Get(p point) rune {
	val, exists := c.points[p]
	if exists {
		return val
	}
	return ' '
}

func (c *MapCanvas) Set(p point, r rune) bool {
	if !c.Contains(p) {
		return false
	}
	c.points[p] = r
	return true
}
