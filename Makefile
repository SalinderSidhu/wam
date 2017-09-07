# Configurable variables
BINARY=wam
GOARCH=amd64
VERSION=0.0.8

# Setup required HASH for the binary and LDFLAGS for the compiler
HASH=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags="-X main.Name=$(BINARY) -X main.Version=$(VERSION) -X main.Hash=$(HASH)"

# Build all binaries
all: clean linux mac win

# Builds the project for Linux based systems
linux:
	GOOS=linux GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)-linux-$(GOARCH)

# Builds the project for Darwin based MAC OS X systems
mac:
	GOOS=darwin GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)-darwin-$(GOARCH)

# Builds the project for Windows based systems
win:
	GOOS=windows GOARCH=$(GOARCH) go build $(LDFLAGS) -o bin/$(BINARY)-windows-$(GOARCH).exe

# Cleans the project: deletes binaries
clean:
	rm -rf ./bin

.PHONY: linux mac win clean
