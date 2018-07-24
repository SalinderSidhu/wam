BINARY=wam
OUTPUT=./bin
SEMVER=0.5.0

VER:=$(SEMVER) HASH_$(shell git rev-parse HEAD | cut -c 1-10 | tr a-z A-Z)
OWNER=$(shell git log -1 --pretty=format:'%an')
EMAIL=$(shell git log -1 --pretty=format:'%ae')
LDFLAGS='-X "main.Owner=$(OWNER)" -X "main.Email=$(EMAIL)" -X "main.Ver=$(VER)"'

OS=$(shell uname -s)
DEPVER:=$(shell dep version 2>/dev/null)

define build
	go build -o $(OUTPUT)/$(BINARY)$(1) -ldflags $(LDFLAGS) *.go
endef

# Build the project
all: clean
ifeq ($(shell uname -s), Darwin)
	$(call build)
else
	$(call build,.exe)
endif

# Setup required dependencies for the project
setup:
ifndef DEPVER
ifeq (OS, Darwin)
	brew install dep
else
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
endif
	dep ensure -vendor-only
.PHONY: setup

# Clean the build folder
clean:
	rm -rf $(OUTPUT)
.PHONY: clean
