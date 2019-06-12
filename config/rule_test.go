package config

import (
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
     bob+1
`
	rules, err := Load([]byte(yaml))
	assert.NoError(t, err)
	data, err := rules.Rules[0].Do(map[string]interface{}{"bob": 2})
	assert.NoError(t, err)
	assert.Equal(t, 3, data)
}
