package timeline

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeline(t *testing.T) {
	tl := New(3)
	assert.Equal(t, 0, tl.Lenght())
	now := time.Now()
	tl.Set(now.Add(-30*time.Second), 1)
	tl.Set(now.Add(-20*time.Second), 2)
	tl.Set(now.Add(-10*time.Second), 3)
	tl.Set(now, 4)
	assert.Equal(t, 2, tl.First())
	assert.Equal(t, 4, tl.Last())
	assert.Equal(t, 3, tl.Lenght())
	short := tl.Duration(-25 * time.Second)
	assert.Equal(t, 2, short.Lenght())
}
