# Configurable variables
SRC=wam.go
SEMVER=0.0.9

# Required variables for the binary and LDFLAGS for the compiler
VERSION:=$(SEMVER) HASH_$(shell git rev-parse HEAD | cut -c1-10)
AUTHOR=$(shell git log -1 --pretty=format:'%an')
EMAIL=$(shell git log -1 --pretty=format:'%ae')

LDFLAGS='-X "main.Author=$(AUTHOR)" -X "main.Email=$(EMAIL)" -X "main.Version=$(VERSION)"'

# Build all binary
all: clean 
	go build -ldflags $(LDFLAGS) $(SRC)

# Cleans the project
clean:
	go clean

.PHONY: clean
