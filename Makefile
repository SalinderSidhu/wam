SEMVER=0.5.1

OWNER=$(shell git log -1 --pretty=format:'%an')
EMAIL=$(shell git log -1 --pretty=format:'%ae')
VER=$(SEMVER) HASH_$(shell git rev-parse HEAD | cut -c 1-10 | tr a-z A-Z)
LDFLAGS='-X "main.Owner=$(OWNER)" -X "main.Email=$(EMAIL)" -X "main.Ver=$(VER)"'

all: clean setup install

# Build the project
build: clean
	go build -ldflags $(LDFLAGS) *.go
.PHONY: build

# Install the project to the go path
install: clean
	go install -ldflags $(LDFLAGS) *.go
.PHONY: install

# Setup required dependencies for the project
setup:
ifeq (, $(shell dep version 2>/dev/null))
ifeq ($(shell uname -s), Darwin)
	brew install dep
else
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep init
endif
	dep ensure -vendor-only
.PHONY: setup

# Remove the compiled project
clean:
	go clean
	rm -f wam.json
.PHONY: clean
