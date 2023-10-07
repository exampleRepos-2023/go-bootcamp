# Common Go Modules Commands

## `$ go help mod`
- List the available go mod commands.
  ```
  $ go help mod
  
  Usage:
          go mod <command> [arguments]
  
  The commands are:
          download    download modules to local cache
          edit        edit go.mod from tools or scripts
          graph       print module requirement graph
          init        initialize new module in current directory
          tidy        add missing and remove unused modules
          vendor      make vendored copy of dependencies
          verify      verify dependencies have expected content
          why         explain why packages or modules are needed
  ```

## `$ go mod init`
- Initializes new module in current directory. 
- Creates the `go.mod` file stating module path & Go version
  ```
  module gitlab.com/golangdojo/bootcamp
    
  go 1.18
  ```


## `$ go get`
Updates `go.mod` and `go.sum` for dependencies. 

- Master 

  `go get github.com/gorilla/mux`
- Branch

  `go get github.com/gorilla/mux@development` 
- Commit hash

  `go get github.com/gorilla/mux@adcd1234` 
- Version

  `go get github.com/gorilla/mux@v1.0.0`
- Update

  `go get -u github.com/gorilla/mux`

  ```
  module gitlab.com/golangdojo/bootcamp
    
  go 1.18
    
  require github.com/gorilla/mux v1.0.0
  ```

## `$ go mod tidy`
- Adds missing and removes unused modules.
- Goes through project files to figure out what to add and remove.
- Creates the go.sum file if needed
  ```
  gitlab.com/golangdojo/bootcamp v0.0.0-20220313194012-d4368af54e69 h1:blxsZAlToSadbu3REsPW+5+HP6/dhyoiQwLMT+QcGqM=
  gitlab.com/golangdojo/bootcamp v0.0.0-20220313194012-d4368af54e69/go.mod h1:IOuGpWCqA123TL+Gxh+xMy1klqokw7dQW45K/4qaDxU=
  github.com/gorilla/mux v1.0.0 h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=
  github.com/gorilla/mux v1.0.0/go.mod h1:DVbg23sWSpFRCP0SfiEN6jmj59UnW/n46BH5rLB71So=
  ```

## `$ go mod edit -GO=1.19`
- Edits Go version to 1.19.
- You may use flags other than -GO= and flag values other than 1.19.
- Updates the go.mod file for the edit.
  ```
  module gitlab.com/golangdojo/bootcamp
    
  go 1.19
  ```

## `$ go mod graph`
- Prints dependency graph.
- You may want to use a third-party library for prettier graphs instead of just texts.
  ```
  $ go mod graph
  gitlab.com/golangdojo/bootcamp github.com/gorilla/mux v1.0.0-20220313194012-d4368af54e69
  ```

## `$ go mod why fmt`
- Explains why fmt is needed in the project.
- Prints the shortest path to the package in question, or nothing if there isn't one.
  ```

  $ go mod why fmt
  # fmt
  gitlab.com/golangdojo/bootcamp/2intermediate/1packagemanagement/gomodcommands
  fmt

  ```

## `$ go mod download`
- Downloads remote dependencies to local cache.
- Uses the go.mod and go.sum file to figure out what to download.

## `$ go mod verify`
- Verifies dependencies downloaded locally are unmodified remotely.
- Prints "all modules verified" if unmodified, or exits with a non-zero status.

## `$ go mod vendor`
- Makes a vendored copy of the dependencies into your own project.
- Use this if you're paranoid the original vendor may update/remove your dependencies.
