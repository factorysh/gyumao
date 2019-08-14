package probes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	f, err := NewFileFromYAML([]byte(`
---
- a
- b
- c

`))
	assert.NoError(t, err)
	assert.True(t, f.Exist("a"))
	assert.False(t, f.Exist("d"))
	assert.Equal(t, []string{"d"}, f.Unknown())
}
