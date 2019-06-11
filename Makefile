build: bin
	dep ensure
	go build -o bin/gyumao

bin:
	mkdir -p bin