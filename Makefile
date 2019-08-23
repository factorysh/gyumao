.PHONY: plugins

all: plugins build

build: bin
	dep ensure
	go build -o bin/gyumao

bin:
	mkdir -p bin

test: plugins
	go test -cover \
		github.com/factorysh/gyumao/deadman \
		github.com/factorysh/gyumao/evaluator/expr \
		github.com/factorysh/gyumao/plugin \
		github.com/factorysh/gyumao/probes \
		github.com/factorysh/gyumao/states \
		github.com/factorysh/gyumao/statesbolt \
		github.com/factorysh/gyumao/point \
		github.com/factorysh/gyumao/rule \
		github.com/factorysh/gyumao/timeline

tests: | test build

workinghours: _plugins
	cd plugins/workinghours && make
	cp plugins/workinghours/workinghours _plugins/ 

plugins: workinghours

_plugins:
	mkdir -p _plugins