# WAM
> A World of Warcraft Addon Manager CLI tool built using Go.

# Table of Contents

* [Getting Started](#getting-started)
    * [Requirements](#requirements)
    * [Building](#Building)
* [Credits](#credits)
* [License](#license)

# Getting Started

## Requirements
* [`The Go Programming Language`](https://golang.org/)
> Version 1.8 or higher is recommended.

This application requires the following go packages:

Package | Command
-- | --
cli | `go get github.com/urfave/cli`
color | `go get github.com/fatih/color`
osext | `go get github.com/kardianos/osext`
goquery | `go get github.com/PuerkitoBio/goquery`
tablewriter | `go get github.com/olekukonko/tablewriter`

## Building

Install the above Go packages and type `make` to compile the binaries. To build platform specific binaries type `make linux` for linux and `make darwin` for Mac OS X.

### Windows

If you're using Windows first download and install [Git for Windows](https://git-for-windows.github.io/). Next, visit [this Gist](https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058) to install the **make** tool for Git Bash. Finally, open Git Bash and type `make windows`.

*Note: The Makefile uses standard Go commands. For more info, please visit the [Offical Documentation](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) on `go build`.*

# Credits

CLI design inspired from [wow-am](https://www.npmjs.com/package/wow-am)

# License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](/LICENSE.md)

**THIS APPLICATION IS IN NO WAY AFFILIATED OR ENDORSED BY BLIZZARD ENTERTAINMENT**
