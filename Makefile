.DEFAULT_GOAL := build

.PHONY:fmt vet build strip
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o hello

clean:
	go clean

strip: build
	strip -s hello
