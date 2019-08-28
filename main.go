package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/factorysh/gyumao/config"
	_gyumao "github.com/factorysh/gyumao/gyumao"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {
	filenameHook := filename.NewHook()
	log.AddHook(filenameHook)
	log.SetLevel(log.DebugLevel)
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
	cfg.Default()

	gyumao, err := _gyumao.New(&cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(gyumao.Serve())
}
