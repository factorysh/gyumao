.PHONY: plugins

all: plugins build

build: bin
	dep ensure
	go build -o bin/gyumao

bin:
	mkdir -p bin

test: plugins
	go test -v \
		github.com/factorysh/gyumao/timeline \
		github.com/factorysh/gyumao/store \
		github.com/factorysh/gyumao/config \
		github.com/factorysh/gyumao/rule \
		github.com/factorysh/gyumao/plugin

workinghours: _plugins
	cd plugins/workinghours && make
	cp plugins/workinghours/workinghours _plugins/ 

plugins: workinghours

_plugins:
	mkdir -p _plugins