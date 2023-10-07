# Package Structure

## Access scopes
- Block
  ```
  import "fmt"
  
  func main() {
    {
      var pi = 3.14
      fmt.Println(pi) // Compiles 
    }
    fmt.Println(pi) // Doesn't compile
  }
  ```

- Unexported - accessible within the same package

  - Same file
    ```
    import "fmt"
  
    var pi = 3.14
  
    func main() {
      {
        fmt.Println(pi) // Compiles 
      }
      fmt.Println(pi) // Compiles
    }
    ```
  - Different file. See "unexported" directory

- Exported - accessible to other packages

  - See "exported" directory

  - Import path
    - Standard packages
    
      `import "fmt"` or `import "math"`
    - Non-standard packages
    
      `import "gitlab.com/golangdojo/bootcamp/2intermediate/1packagemanagement/packagestructure/exported/math"`
      - Host: `gitlab.com`
      - User/company: `golangdojo`
      - Project: `bootcamp`
      - Path to package: `/2intermediate/1packagemanagement/packagestructure/exported/`
      - Package: `math`

