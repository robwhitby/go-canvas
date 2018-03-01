package canvas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMapCanvas(t *testing.T) {
	c := NewMapCanvas(4, 5)
	assert.Equal(t, 4, c.Width())
	assert.Equal(t, 5, c.Height())
}

func TestMapCanvas_Contains(t *testing.T) {
	c := NewMapCanvas(4, 5)
	assert.True(t, c.Contains(0, 0))
	assert.True(t, c.Contains(3, 4))
	assert.False(t, c.Contains(4, 3))
	assert.False(t, c.Contains(4, 5))
}

func TestMapCanvas_Get_EmptyPoint(t *testing.T) {
	c := NewMapCanvas(4, 5)
	assert.Equal(t, ' ', c.Get(1, 2))
}

func TestMapCanvas_Set_InBounds(t *testing.T) {
	c := NewMapCanvas(4, 5)
	assert.True(t, c.Set(1, 2, 'x'))
	assert.Equal(t, 'x', c.Get(1, 2))
}

func TestMapCanvas_Set_OutOfBounds(t *testing.T) {
	c := NewMapCanvas(4, 5)
	assert.False(t, c.Set(4, 5, 'x'))
}
