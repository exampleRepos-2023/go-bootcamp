# Go Workspaces

```
workspace
    |_ project1
           |_ go.mod 
              module golangdojo.com/golangdojo/project1
    |_ project2
           |_ go.mod 
              module golangdojo.com/golangdojo/project2
    |_ go.work
       go 1.18
       use ./project1
       use ./project2
```

No need to `cd` into the directories - Multi-module workspaces

`$ go run golangdojo.com/golangdojo/project1`

`$ go run golangdojo.com/golangdojo/project2`

(Also, found a typo)

