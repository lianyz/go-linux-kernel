.DEFAULT: all

.PHONY: all
all: build run

.PHONY: build
build:
	go build

.PHONY: run
run:
	./go-linux-kernel run /bin/bash


.PHONY: run.echo
run.echo:
	./go-linux-kernel run echo hello

