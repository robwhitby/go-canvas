package go_canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMapCanvas(t *testing.T) {
	c := NewCanvas(4, 5)
	assert.Equal(t, 4, c.Width())
	assert.Equal(t, 5, c.Height())
}

func TestMapCanvas_Contains(t *testing.T) {
	c := NewCanvas(4, 5)

	good := []point{{1, 1}, {4, 5}, {3, 4}}
	for _, p := range good {
		assert.True(t, c.Contains(p.x, p.y))
	}

	bad := []point{{0, 0}, {0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {5, 4}}
	for _, p := range bad {
		assert.False(t, c.Contains(p.x, p.y))
	}
}

func TestMapCanvas_Get_EmptyPoint(t *testing.T) {
	c := NewCanvas(4, 5)
	assert.Equal(t, ' ', c.Get(1, 2))
}

func TestMapCanvas_Set_InBounds(t *testing.T) {
	c := NewCanvas(4, 5)
	c.Set(1, 2, 'x')
	assert.Equal(t, 'x', c.Get(1, 2))
}

func TestMapCanvas_Recreate(t *testing.T) {
	c := NewCanvas(4, 5)
	c.Set(1, 1, 'x')
	c.Recreate(1, 2)
	assert.Equal(t, 1, c.Width())
	assert.Equal(t, 2, c.Height())
	assert.Equal(t, ' ', c.Get(1, 1))
}
