build: vendor
	go build \
		-ldflags "-X github.com/factorysh/gyumao/plugins/workinghours/main.version=`git rev-parse HEAD`"\
		-o workinghours

vendor:
	dep ensure