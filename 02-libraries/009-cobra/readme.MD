## Commands
`go mod init go-3-bookstore`
- initialized a go.mod file
- Describes the module: with name/module path go version used in the program
- The module path is also the import path (e.g. import "github/nana/booking-app")
- Creates a new module
- Module path can correspond to a repository you plan to publish your module.

`go get github.com/spf13/cobra` --> to get the cobra library

`go install github.com/spf13/cobra/cobra-cli` --> to get a help tool that creates the boilerplate code for us

`cobra-cli add <name>` --> to add a file, in this case could be `code`

- Notes
   - 2 options --> using arguments, and using flags

`go run <file name>` compiles and runs the code