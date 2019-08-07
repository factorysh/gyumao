package plugin

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	fmt.Println("Host pid", os.Getpid())
	p, err := getPlugin("../plugins/workinghours/workinghours")
	assert.NoError(t, err)
	err = p.Setup(map[string]interface{}{"age": 42, "name": "Robert"})
	assert.NoError(t, err)
	m, err := p.Meta()
	assert.NoError(t, err)
	fmt.Println("meta", m)
	assert.Equal(t, "hours", m.Class)
}
