package deadman

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCircular(t *testing.T) {
	g := New(3, []string{"a", "b", "c", "d"}, time.Hour)
	assert.Equal(t, 0, g.rank)
	g.Tick()
	assert.Equal(t, 1, g.rank)
	g.Tick()
	assert.Equal(t, 2, g.rank)
	g.Tick()
	assert.Equal(t, 0, g.rank)
}

func TestPrevious(t *testing.T) {
	g := New(3, []string{"a", "b", "c", "d"}, time.Hour)
	g.Tick()
	assert.Equal(t, 1, g.rank)
	i := g.previous(0)
	assert.Equal(t, 1, i)
	i = g.previous(1)
	assert.Equal(t, 0, i)
	i = g.previous(2)
	assert.Equal(t, 2, i)
}

func TestCrunch(t *testing.T) {
	g := New(3, []string{"a", "b", "c", "d"}, time.Hour)
	g.Current().Alive("a").Alive("b")
	g.Tick()
	g.Current().Alive("b").Alive("c")
	c := g.Crunch(1)
	fmt.Println(c.bitset)
	i := c.DeadIterator()
	cpt := 0
	fmt.Println(i.Next())
	for n, ok := i.Next(); ok; n, ok = i.Next() {
		cpt++
		assert.Equal(t, "d", n)
	}
	assert.Equal(t, 1, cpt)
}
