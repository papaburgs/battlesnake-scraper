.PHONY: deploy build run clean gitc

run: gitc
	go run *.go

build: gitc
	go build *.go

test:
	go test ../../internal/pkg/snake
	go test .

gitc:
	git rev-parse --short HEAD >gitc.txt


