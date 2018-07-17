# Configurable variables
SEMVER=0.5.0
FOLDER=bin
BINARY=wam.exe

# Required variables for the binary and LDFLAGS for the compiler
VERSION:=$(SEMVER) dev_hash_$(shell git rev-parse HEAD | cut -c1-20)
AUTHOR=$(shell git log -1 --pretty=format:'%an')
EMAIL=$(shell git log -1 --pretty=format:'%ae')

LDFLAGS='-X "main.Author=$(AUTHOR)" -X "main.Email=$(EMAIL)" -X "main.Version=$(VERSION)"'

# Build the binary
all: clean 
	go build -o $(FOLDER)/$(BINARY) -ldflags $(LDFLAGS) *.go

# Clean the project
clean:
	go clean
	rm -rf ./$(FOLDER)

.PHONY: clean
