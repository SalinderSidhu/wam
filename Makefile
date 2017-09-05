# Configurable variables
BINARY=wam
GOARCH=amd64
VERSION=0.0.8

# Setup required HASH for the binary and LDFLAGS for the compiler
HASH=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="-X main.Name=$(BINARY) -X main.Version=$(VERSION) -X main.Hash=$(HASH)"

# Build the entire project
all: clean linux darwin windows

# Builds the project
linux:
	GOOS=linux GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)

# Builds the project
darwin:
	GOOS=darwin GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)

# Builds the project
windows:
	GOOS=windows GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY).exe

# Cleans the project: deletes binaries
clean:
	rm -rf ./bin

.PHONY: linux darwin windows clean
