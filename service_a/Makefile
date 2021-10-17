.PHONY: \
	test \
	install-lint \
	lint \

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