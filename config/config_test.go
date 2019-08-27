package config

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {
	raw, err := ioutil.ReadFile("../gyumao.yml")
	assert.NoError(t, err)
	var cfg Config
	err = yaml.Unmarshal(raw, &cfg)
	assert.NoError(t, err)
	cfg.Default()
}
