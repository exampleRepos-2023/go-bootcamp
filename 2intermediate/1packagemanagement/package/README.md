# Package

## What is a package? 
Packages are a grouping unit of code pieces, and they help you organize your code. 

You can put multiple files with shared attributes into a single package.

## Declaring a package

```

// Inside main.go file

package main

```


## Types of packages

- Executable package (main)

    `main.go`

    ```
    package main

    import "fmt"

    func main() {
      fmt.Println("Hello World")
    }
    ```

- Non-executable packages (non-main)

  `math.go`

  ```
  package math
  
  func add(a, b int) int {
    return a + b
  }
  
  func subtract(a, b int) int {
    return a - b
  }
  ```


## Package naming

- Short. No camel/snake case - prompts for long package names. 

  `math` instead of `numericCalculationConstantsAndFunctions`

- Descriptive. No util/common/etc. - doesn't say much about the package.

- Same as your directory name. Don't declare multiple packages in multiple files under the same directory

  ```
  app
    |_ main
         |_ main.go
                |_ 
                   // Inside the file
                   package main
                   ...
    |_ math
         |_ math.go
                |_ 
                   // Inside the file
                   package math
                   ...
  ```
  instead of
  ```
  app
    |_ main
         |_ main.go
         |      |_ 
                   // Inside the file
         |         package main
         |         ...
         |_ math.go
                |_ 
                   // Inside the file
                   package math
                   ...
  ```


- Has a file with the same name to serve as the primary file for the package

  ```
  app
    |_ math
         |_ math.go
                |_ 
                   // Inside the file
                   package math
                   func add(a, b int) int {return a + b}
                   func subtract(a, b int) int {return a - b}
  ```
  instead of

  ```
  app
    |_ math
         |_ add.go
         |      |_ 
         |         // Inside the file
         |         package math
         |         func add(a, b int) int {return a + b}
         |         
         |_ subtract.go
                |_
                   // Inside the file
                   package math
                   func subtract(a, b int) int {return a - b}
  ```


- Small and many. Less nested packages.

  Don't do following.

  ```
  app
    |_ math
         |_ int
              |_ positive
                   |_ add
                   |_ subtract
              |_ negative
                   |_ add
                   |_ subtract
              |_ zero
                   |_ add
                   |_ subtract
  ```

- ...

## 