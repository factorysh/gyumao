package timeline

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeline(t *testing.T) {
	tl := NewTimeline(3)
	assert.Equal(t, 0, tl.Lenght())
	tl.Set(time.Now(), 1)
	tl.Set(time.Now(), 2)
	tl.Set(time.Now(), 3)
	tl.Set(time.Now(), 4)
	assert.Equal(t, 2, tl.First())
	assert.Equal(t, 4, tl.Last())
	assert.Equal(t, 3, tl.Lenght())
}
