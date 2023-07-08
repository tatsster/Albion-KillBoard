.PHONY: default

default: build;

build:
	go build -o out/killbot ./cmd/main.go