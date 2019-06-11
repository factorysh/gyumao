package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	yaml := `
rules:
 - name: cpu
   keys:
    - hostname
   expr: >
     1+1
`
	rules, err := Load([]byte(yaml))
	assert.NoError(t, err)
	fmt.Println(rules)
}
