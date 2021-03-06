# WAM
[![Go Report Card](https://goreportcard.com/badge/github.com/salindersidhu/wam)](https://goreportcard.com/report/github.com/salindersidhu/wam)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](/LICENSE.md)

A World of Warcraft addon manager CLI tool built using Go.

# Table of Contents
* [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Building](#building)
    * [Installing](#installing)
* [Contributing](#contributing)

# Getting Started

## Requirements
* [`The Go Programming Language`](https://golang.org/)
* [`Make tool for Git Bash on Windows`](https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058)

## Building
If you're using Windows, please download and install the **make** tool for Git Bash.

Install required GO packages and update dependencies for this project:
```bash
# Installs the dep tool when executed for the first time
make setup
```

Compile the project for your OS and output the binary to your current working folder:
```bash
make build
```

Clean the compiled project:
```bash
make clean
```

## Installing

Install the project as a standalone tool to the go path binary folder:
```bash
make install
```

# Contributing
Wam welcomes contributions from anyone and everyone. Please see our [contributing guide](/CONTRIBUTING.md) for more info.

**THIS APPLICATION IS IN NO WAY AFFILIATED OR ENDORSED BY BLIZZARD ENTERTAINMENT**
