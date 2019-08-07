package plugin

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	plugins := NewPlugins()
	err := plugins.register("../plugins/workinghours/workinghours", nil)
	assert.NoError(t, err)
	w := plugins.HoursPlugins["workinghours"]
	tags, err := w.Time(time.Date(2018, time.November, 11, 11, 11, 0, 0, time.UTC))
	assert.NoError(t, err)
	fmt.Println(tags)
	assert.Equal(t, "working", tags["hours"])
}
