package deadman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadman(t *testing.T) {
	d := New(3)
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
}
