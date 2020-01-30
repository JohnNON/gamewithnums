.PHONY: build
build:
	go build -v ./cmd/gameserver

.PHONY: test
test:
	go test -v -race -temeout 30s ./...

.DEFAULT_GOAL := build