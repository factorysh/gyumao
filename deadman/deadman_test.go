package deadman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadman(t *testing.T) {
	d := NewDeadRegistry([]string{"a", "b", "c"})
	assert.False(t, d.bitset.Any())
	d.Alive("a")
	d.Alive("c")
	fmt.Println(d.bitset)
	i := d.DeadIterator()
	v, ok := i.Next()
	assert.True(t, ok)
	assert.Equal(t, "b", v)
	v, ok = i.Next()
	fmt.Println(v)
	assert.False(t, ok)

	assert.True(t, d.bitset.Any())
	d.Reset()
	assert.False(t, d.bitset.Any())
}

func TestIterator(t *testing.T) {
	d := NewDeadRegistry([]string{"a", "b", "c", "d"})
	d.Alive("b").Alive("c")
	i := d.DeadIterator()
	cpt := 0
	for n, ok := i.Next(); ok; n, ok = i.Next() {
		fmt.Println("n", n)
		cpt++
	}
}

func TestGhost(t *testing.T) {
	d := NewDeadRegistry([]string{"a", "b", "c", "d"})
	d2 := d.Ghost()
	assert.Equal(t, "b", d2.keys[1])
	d.Alive("a")
	assert.True(t, d.bitset.Test(0))
	assert.False(t, d2.bitset.Test(0))
}

func TestEnlarge(t *testing.T) {
	d := NewDeadRegistry([]string{"a", "b"})
	err := d.Add("c", "d")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c", "d"}, d.keys)
	r, ok := d.keysRank.Get("c")
	assert.True(t, ok)
	assert.Equal(t, r, uint(2))
}

func TestRemove(t *testing.T) {
	d := NewDeadRegistry([]string{"a", "b", "c", "d"})
	d.Alive("b")
	err := d.Remove("c")
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "d"}, d.keys)
	r, ok := d.keysRank.Get("b")
	assert.True(t, ok)
	assert.Equal(t, uint(1), r)
	assert.True(t, d.bitset.Test(r.(uint)))
	assert.False(t, d.bitset.Test(uint(0)))
	assert.False(t, d.bitset.Test(uint(2)))
}
