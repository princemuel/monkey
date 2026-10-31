
BINARY  := monkey
BIN_DIR := bin
CMD     := ./cmd/monkey

.DEFAULT_GOAL := build

.PHONY: fmt vet test build run strip clean

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

# Entry point lives in cmd/monkey; the binary lands in bin/ (gitignored).
build: vet
	go build -o $(BIN_DIR)/$(BINARY) $(CMD)

run: build
	./$(BIN_DIR)/$(BINARY)

# Smaller release binary: -s drops the symbol table, -w drops DWARF debug info.
strip: vet
	go build -trimpath -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY) $(CMD)

clean:
	rm -rf $(BIN_DIR)
