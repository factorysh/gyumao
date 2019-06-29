package deadman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircular(t *testing.T) {
	g := NewGenealogy(3, 42)
	assert.Equal(t, 0, g.rank)
	g.Tick()
	assert.Equal(t, 1, g.rank)
	g.Tick()
	assert.Equal(t, 2, g.rank)
	g.Tick()
	assert.Equal(t, 0, g.rank)
}

func TestPrevious(t *testing.T) {
	g := NewGenealogy(3, 42)
	g.Tick()
	assert.Equal(t, 1, g.rank)
	i := g.previous(0)
	assert.Equal(t, 1, i)
	i = g.previous(1)
	assert.Equal(t, 0, i)
	i = g.previous(2)
	assert.Equal(t, 2, i)
}
