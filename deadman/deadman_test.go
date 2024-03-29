package deadman

import (
	"fmt"
	"sync"
	"testing"

	"math/rand"

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

func TestConcurrency(t *testing.T) {
	size := 1000
	parralel := 8
	queue := make(chan string)

	keys := make([]string, size)
	for i := 0; i < size; i++ {
		keys[i] = fmt.Sprintf("%d", i)
	}
	d := NewDeadRegistry(keys)

	w := sync.WaitGroup{}
	w.Add(size)

	go func() {
		for _, v := range rand.Perm(size) {
			fmt.Println("<-", v)
			queue <- fmt.Sprintf("%d", v)
		}
	}()

	for i := 0; i < parralel; i++ {
		go func() {
			for {
				v := <-queue
				fmt.Println("   ", v)
				d.Alive(v)
				w.Done()
			}
		}()
	}

	w.Wait()
}
