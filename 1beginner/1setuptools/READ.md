# Tools & Setup

## Intro to Go
https://go.dev/
- Created in 2007, v.1.0 & open-sourced in 2012
- Initally aimed to replaced C
- Migration target from JavaScript, Python, Ruby, Java
- Static typed, compiled
- 3/4 OOP
- Simple* (error handling, generics, packages)
- Concurrency (goroutine, channels)

## Installing Go
1. Download runtime
   - https://go.dev/dl/
   - `wget`
2. Install runtime
   - Unzip `tar -xzvf`
   - Go runtime folder `/usr/local/go`
3. Create workspace folder
   - bin - binaries/executables of your projects
   - pkg - imported packages
   - src - your project source code goes here
   - ```
     /HOME
       /workspace
         /bin
         /pkg
         /src
4. Set environment variables 
   - PATH - creating shortcuts for executables you may use frequently
   - GOROOT - where go runtime executables live
   - GOPATH -  where local go project executables live
   - `$` indicates it's an existing variable
   - ```
     # User settings
     export GOROOT=/usr/local/go
     export GOPATH=$HOME/goworkspace
     export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
5. Check with `go version`

## Installing VS Code
Most popular general code editor

Pros
1. Free
2. Versatile
3. Popular

Con
1. Requires extensions
2. Requires additional packages for those extensions
3. Unexpected hiccups

Installation
1. Download via
   https://code.visualstudio.com/
2. Install the "Go" extension
3. Install packages for the "Go" extension
4. Hello World


## Installing Goland
Industry standard IDE specifically for Go development

Pros
1. No extensions needed
2. No additional packages needed 
3. Works

Con
1. Not free
2. Fewer extension selections
3. Doesn't work well with other languages

Installation
1. Download via https://www.jetbrains.com/go/
2. Select Go runtime
3. Hello World

## Hello World
```
package main

import "fmt"

func main() {
   fmt.Println("Hello World!")
}

```

## Basic Go Commands
- go version
  - Prints out the version of the Go runtime installed
- go env
  - Prints out a list of Go environment variables 
- go fmt
  - Formats given Go files
- go run
  - Executes given a Go file/program
- go build
  - Builds an executable/binary given a Go file/program