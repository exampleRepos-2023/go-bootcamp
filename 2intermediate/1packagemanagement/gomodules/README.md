# Go Modules

## What is Go Modules?
Go Modules is native dependency management tool.

## What is a dependency?
A resource that the project depends on. In Go, package is a resource unit.

## Why Go Modules
Primary functions
- Dependency Requirements
  - What these dependencies are
  - The path to get them
- Dependency Versioning
  - Specifying versions of the required dependencies

Other benefits
- Initialize a Go project outside $GOPATH/src
- De facto solution unlike previous third-party ones like `dep` or `govendor`
- ...

## How Go Modules works
- Initialize a Go project/ making the project a module

  - `go mod init [project path]`

    `go mod init gitlab.com/golangdojo/bootcamp`

- Pull an external project or package

  - Project

    `go get golang.org/x/crypto`
  - Package

    `go get golang.org/x/crypto/bcrypt`

  - (Optional) Specify a version

    `go get golang.org/x/crypto@v0.0.0`

- Use the dependency
  - Import the package in a .go file
    
    `import "golang.org/x/crypto/bcrypt"`

- Location
  - Dependency is downloaded in
    `$GOPATH/pkg/mod`
- Enabled by 
  - `go.mod`
  - `go.sum`

