package plugin

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	fmt.Println("Host pid", os.Getpid())
	m, err := meta("../plugins/workinghours/workinghours")
	assert.NoError(t, err)
	fmt.Println("meta", m)
	assert.Equal(t, "hours", m.Class)
}
