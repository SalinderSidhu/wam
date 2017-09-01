# Name of the binary output
BINARY=wam

# VERSION value to pass into the binary 
VERSION=0.0.6

# Setup the -ldflags option for go build, interpolate the variable values
LDFLAGS=-ldflags "-X main.Name=$(BINARY) -X main.Version=$(VERSION)"

# Default target
.DEFAULT_GOAL: build

# Builds the project
build:
	go build $(LDFLAGS) -o bin/$(BINARY)

# Cleans the project: deletes binaries
clean:
	rm -rf bin/$(BINARY)

.PHONY: clean install
