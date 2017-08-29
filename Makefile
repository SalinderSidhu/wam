# Name of the binary output
BINARY=wam

# VERSON value to pass into the binary 
VERSION=0.0.5

# Setup the -ldflags option for go build, interpolate the variable values
LDFLAGS=-ldflags "-X main.Name=${BINARY} -X main.Version=${VERSION}"

# Default target
.DEFAULT_GOAL" ${BINARY}

# Builds the project
${BINARY}:
	go build ${LDFLAGS} -o {$BINARY}

# Installs the project: copies binaries
install:
	go install ${LDFLAGS} -o {$BINARY}

# Cleans the project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
