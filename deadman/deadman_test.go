package deadman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadman(t *testing.T) {
	d := New(3)
	assert.False(t, d.bitset.Any())
	d.Alive(0)
	d.Alive(2)
	fmt.Println(d.bitset)
	i := d.DeadIterator()
	v, ok := i.Next()
	assert.True(t, ok)
	assert.Equal(t, uint(1), v)
	v, ok = i.Next()
	fmt.Println(v)
	assert.False(t, ok)

	assert.True(t, d.bitset.Any())
	d.Reset()
	assert.False(t, d.bitset.Any())
}
