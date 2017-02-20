VERSION := $(shell git describe --always --tags --dirty --long)
BUILD_DATE := $(shell date -u +"%Y.%m.%d-%H:%M:%S")
GO_VERSION := $(shell go version | cut -c 12- | sed -e 's/ /-/g')

BUILD_FLAGS := -v -i -ldflags "-s -w -X github.com/TheHippo/snc.version=${VERSION} -X github.com/TheHippo/snc.date=${BUILD_DATE} -X github.com/TheHippo/snc.goVersion=${GO_VERSION}"

OUTPUT := "snc"

all: snc

test: *.go **/*.go
	go test -v github.com/TheHippo/snc/...

lint: *.go **/*.go
	golint github.com/TheHippo/snc/...

vet: *.go **/*.go
	go vet github.com/TheHippo/snc/...

snc: lint vet test *.go **/*.go
	go build -o ${OUTPUT} ${BUILD_FLAGS} github.com/TheHippo/snc

clean:
	@-rm ${OUTPUT}
