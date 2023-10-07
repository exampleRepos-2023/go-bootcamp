# Package Management


## Types of packages
- Standard packages 
  - Lives in $GOROOT/src (Go runtime)
  - `import "fmt"`
- Internal packages
  - Lives in $GOPATH/src/... (Go workspace)
  - `import "gitlab.com/golangdojo/bootcamp/.../math"`
- External packages
  - Developed and hosted (ie, gitlab.com) externally
  - Imported and used just like an internal package afterwards
  - Go Modules
