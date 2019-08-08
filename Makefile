.PHONY: plugins

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

workinghours:
	make -f plugins/workinghours/Makefile

plugins: workinghours