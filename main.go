package main

import (
	"os"

	"io/ioutil"

	"github.com/factorysh/gyumao/config"
	_gyumao "github.com/factorysh/gyumao/gyumao"
	"gopkg.in/yaml.v2"
)

func main() {
	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		configPath = "/etc/gyumao.yml"
	}
	var cfg config.Config
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(raw, &cfg)
	if err != nil {
		panic(err)
	}

	gyumao := _gyumao.New()
	err = gyumao.Plugins.RegisterAll(cfg.PluginFolder, cfg.Plugins)
	if err != nil {
		panic(err)
	}
}
