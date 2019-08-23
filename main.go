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
	cfg.Default()

	// STATE_TEST
	// router := mux.NewRouter()
	// router.HandleFunc("/environment/{collection}/{id}", statesrest.StatesRESTHandlerId)
	// router.HandleFunc("/environment/{collection}/{id}/{key}", statesrest.StatesRESTHandlerKey)
	// http.ListenAndServe("127.0.0.1:8080", router)

	gyumao, err := _gyumao.New(&cfg)
	if err != nil {
		panic(err)
	}
	gyumao.Serve()
}
