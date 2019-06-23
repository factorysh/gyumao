build: bin
	dep ensure
	go build -o bin/gyumao

bin:
	mkdir -p bin

test:
	go test -v \
		gitlab.bearstech.com/factory/gyumao/timeline \
		gitlab.bearstech.com/factory/gyumao/store \
		gitlab.bearstech.com/factory/gyumao/config \
		gitlab.bearstech.com/factory/gyumao/rule
