.DEFAULT: all

.PHONY: all
all: build run

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

.PHONY: tools
tools:
	apt install memtester

.PHONY: run
run:
	./go-linux-kernel run /bin/bash

.PHONY: test
test:
	memtester 100M 1

.PHONY: run.echo
run.echo:
	./go-linux-kernel run echo hello

