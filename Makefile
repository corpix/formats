.DEFAULT_GOAL = all

version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := formats
package  := github.com/corpix/$(name)

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p $@
	go build -o $@/formats ./formats/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -bench=. -v ./...

.PHONY: lint
lint:
	go vet -v ./...

.PHONY: check
check: lint test

.PHONY: clean
clean:
	git clean -xddff
