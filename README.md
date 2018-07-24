# WAM
[![Go Report Card](https://goreportcard.com/badge/github.com/salindersidhu/wam)](https://goreportcard.com/report/github.com/salindersidhu/wam)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](/LICENSE.md)

A World of Warcraft addon manager CLI tool built using Go.

# Table of Contents
* [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Building](#building)
* [Contributing](#contributing)

# Getting Started

## Requirements
* [`The Go Programming Language`](https://golang.org/)
* [`Make tool for Git Bash on Windows`](https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058)

## Building
If you're using Windows, please download and install the **make** tool for Git Bash.

Install the required GO packages and update dependencies for this project:
```bash
# Installs dep tool when executed for the first time
make setup
```

Compile the project for your OS and output to bin folder:
```bash
make
```

Clean compiled project:
```bash
make clean
```

*Fore more info on Go commands, please visit the [Offical Documentation](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies).*

# Contributing
Wam welcomes contributions from anyone and everyone. Please see our [contributing guide](/CONTRIBUTING.md) for more info.

**THIS APPLICATION IS IN NO WAY AFFILIATED OR ENDORSED BY BLIZZARD ENTERTAINMENT**
