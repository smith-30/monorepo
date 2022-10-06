GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# if tag doesnt exists, show revision
VERSION=$(shell git describe --tags --abbrev=7 --always)
REVISION=$(shell git rev-parse --short HEAD)

TEST_RESULTS=hoge

NAME=service_a

.PHONY: \
	test \
	install-lint \
	lint \
	build \

test:
	go test ./...

install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b build/bin v1.42.1

lint:
	./build/bin/golangci-lint run --tests --disable-all \
		--enable=revive \
		--enable=govet \
		--enable=unused \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=structcheck \
		--enable=goimports \
		--enable=gosimple \
		--enable=bodyclose \
		--enable=errname \
		--enable=ifshort \
		--enable=nilerr \

build:
	go build -trimpath -o build/${GOOS}_${GOARCH}/${NAME} main.go