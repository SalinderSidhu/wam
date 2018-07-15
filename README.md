# WAM
[![Go Report Card](https://goreportcard.com/badge/github.com/salindersidhu/wam)](https://goreportcard.com/report/github.com/salindersidhu/wam)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](/LICENSE.md)

A World of Warcraft addon manager CLI tool built using Go.

# Table of Contents

* [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Building](#building)

# Getting Started

## Requirements
* [`The Go Programming Language`](https://golang.org/)

This application requires the following go packages:

```bash
go get github.com/urfave/cli
go get github.com/fatih/color
go get github.com/PuerkitoBio/goquery
go get github.com/olekukonko/tablewriter
```

## Building
Install the above Go packages and type `make` to compile the binary.

If you're using Windows visit [this Gist](https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058) to install the **make** tool for Git Bash. Then, open Git Bash and type `make`.

*Note: The Makefile uses standard Go commands. For more info, please visit the [Offical Documentation](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) on `go build`.*

**THIS APPLICATION IS IN NO WAY AFFILIATED OR ENDORSED BY BLIZZARD ENTERTAINMENT**
