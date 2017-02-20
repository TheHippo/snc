VERSION := $(shell git describe --always --tags --dirty --long)
BUILD_DATE := $(shell date -u +"%Y.%m.%d-%H:%M:%S")
GO_VERSION := $(shell go version | cut -c 12- | sed -e 's/ /-/g')

BUILD_FLAGS := -v -i -ldflags "-s -w -X main.version=${VERSION} -X main.date=${BUILD_DATE} -X main.goVersion=${GO_VERSION}"

OUTPUT := "snc"

all: snc

test:
	go test -v github.com/TheHippo/snc/...

lint:
	golint github.com/TheHippo/snc/...

vet:
	go vet github.com/TheHippo/snc/...

snc: lint vet test
	go build -o ${OUTPUT} ${BUILD_FLAGS} github.com/TheHippo/snc/cmd/snc

clean:
	@-rm ${OUTPUT}
