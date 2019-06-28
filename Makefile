build: bin
	dep ensure
	go build -o bin/gyumao

bin:
	mkdir -p bin

test:
	go test -v \
		github.com/factorysh/gyumao/timeline \
		github.com/factorysh/gyumao/store \
		github.com/factorysh/gyumao/config \
		github.com/factorysh/gyumao/rule
