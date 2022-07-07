## Go Concepts

This repo includes simple POCs around different topics related to Golang

### Useful Commands

- `go mod init <module-name>`
  - Initializes a go.mod file
  - Describes the module: with name/module path go version used in the program
  - The module path is also the import path (e.g. import "github/gcalvoCR/go-concepts")
  - The module path can correspond to a repository you plan to publish your module.
- `go run <file-name>` compiles and runs the file you want to run
- `go get <repo-url>` to get the specific library
- `go mod download` download the dependencies
- `go run .` runs the whole project by doing

### Testing
- `go test -v ./...` runs all tests
- `go test -v <file-path> -run <function-name>` runs a specific function

## Interesting points
- All code must belong to a `package`
- One of the main advantages of Go is Concurrency
- Concurrency is cheap and easy. It prevents the app from being blocked
- By default code gets executed in a single thread
- Concurrency is very simple, using the keyword `go` starts a new goroutine
- A goroutine is lightweight thread managed by the Go runtime


### Interesting links 
- https://www.youtube.com/watch?v=yyUHQIec83I

