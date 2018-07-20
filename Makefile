BINARY=wam
OUTPUT=bin
SEMVER?=DEV

DEPVER:=$(shell dep version 2>/dev/null)
VERSION:=$(SEMVER) HASH_$(shell git rev-parse HEAD | cut -c 1-10 | tr a-z A-Z)
AUTHOR=$(shell git log -1 --pretty=format:'%an')
EMAIL=$(shell git log -1 --pretty=format:'%ae')
LDFLAGS='-X "main.Author=$(AUTHOR)" -X "main.Email=$(EMAIL)" -X "main.Version=$(VERSION)"'

define build
	go build -o $(OUTPUT)/$(BINARY)$(1) -ldflags $(LDFLAGS) *.go
endef

define platform_build
	GOOS=$(1) GOARCH=$(2) $(call build,-$(1)-$(2)$(3))
endef

# Build the binary
all: clean
ifeq ($(shell uname -s), Darwin)
	$(call build)
else
	$(call build,.exe)
endif

# Setup required dependencies for the project
setup:
ifdef DEPVER
	dep ensure -vendor-only
else
ifeq ($(shell uname -s), Darwin)
	brew install dep
else
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep ensure -vendor-only
endif
.PHONY: setup

# Generate binaries for multiple platforms
deploy: clean
	$(call platform_build,windows,amd64,.exe)
	$(call platform_build,darwin,amd64)
.PHONY: deploy

# Clean the build folder
clean:
	rm -rf ./$(OUTPUT)
.PHONY: clean
