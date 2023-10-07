# Understanding go.mod & go.sum

## What are they?
Go Modules use go.mod and go.sum

## What go.mod does
Your project's dependencies

```
module gitlab.com/golangdojo/awesomeproject

go 1.18

require gitlab.com/golangdojo/bootcamp v0.0.0-20220313194012-d4368af54e69 // indirect

```

- Module import path
  - `module gitlab.com/golangdojo/awesomeproject`
- Go version
  - `go 1.18`
- Dependency requirement
  - `require gitlab.com/golangdojo/bootcamp`
- Dependency versioning
  - `v0.0.0`
- Dependency hash
  - `-20220313194012-d4368af54e69`
- If the downloaded dependency is in use
  - `// indirect`

## What go.sum does
Dependencies of your project's dependencies (in addition to your project's dependencies, if any)
- Dependency requirement
- Dependency versioning
- Dependency hash

The Go web framework Fiber uses FastHTTP
```
github.com/gofiber/fiber/v2 v2.22.0 h1:+iyKK4ooDH6z0lAHdaWO1AFIB/DZ9AVo6vz8VZIA0EU=
github.com/gofiber/fiber/v2 v2.22.0/go.mod h1:MR1usVH3JHYRyQwMe2eZXRSZHRX38fkV+A7CPB+DlDQ=
github.com/valyala/fasthttp v1.31.0 h1:lrauRLII19afgCs2fnWRJ4M5IkV0lo2FqA61uGkNBfE=
github.com/valyala/fasthttp v1.31.0/go.mod h1:2rsYD01CKFrjjsvFxx75KlEUNpWNBY9JWD3K/7o2Cus=
...
```