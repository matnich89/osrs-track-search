.PHONY: build test clean generate

GO_FILES := $(shell find . -name '*.go')

default: build

generate:
	go generate ./...

build: generate $(GO_FILES)
	go build -o osrs-track-search

test: generate $(GO_FILES)
	go test -v ./...

clean:
	rm osrs-track-search