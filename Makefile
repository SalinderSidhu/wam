# Configurable variables
BINARY=wam
FOLDER=bin
GOARCH=amd64
VERSION=0.0.8

# Required HASH for the binary and LDFLAGS for the compiler
HASH=$(shell git rev-parse HEAD)
LDFLAGS="-X main.Name=$(BINARY) -X main.Version=$(VERSION) -X main.Hash=$(HASH)"

# Function definition for compiling and building Golang projects
define build
	GOOS=$(1) GOARCH=$(2) go build -ldflags=$(LDFLAGS) -o $(FOLDER)/$(BINARY)$(3)
endef

# Build all binaries for each OS
all: clean 
	$(call build,linux,$(GOARCH),-linux-$(GOARCH))
	$(call build,darwin,$(GOARCH),-darwin-$(GOARCH))
	$(call build,windows,$(GOARCH),-windows-$(GOARCH).exe)

# Builds the project for Linux systems
linux:
	$(call build,linux,$(GOARCH))

# Builds the project for Darwin systems
darwin:
	$(call build,darwin,$(GOARCH))

# Builds the project for Windows systems
windows:
	$(call build,windows,$(GOARCH),.exe)

# Cleans the project
clean:
	rm -rf ./$(FOLDER)

.PHONY: linux darwin windows clean
