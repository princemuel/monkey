BINARY := monkey

.DEFAULT_GOAL := build

.PHONY: fmt vet test build run strip clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

# Needs a main.go in the project root (arrives with the REPL chapter).
build: vet
	go build -o $(BINARY) .

run: build
	./$(BINARY)

# Smaller release binary: -s drops the symbol table, -w drops DWARF debug info.
strip: vet
	go build -trimpath -ldflags="-s -w" -o $(BINARY) .

clean:
	go clean
	rm -f $(BINARY)
