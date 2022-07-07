## Commands
`go mod init go-2-crudapi`
- initialized a go.mod file
- Describes the module: with name/module path go version used in the program
- The module path is also the import path (e.g. import "github/nana/booking-app")
- Creates a new module
- Module path can correspond to a repository you plan to publish your module.

`go get "github.com/gorilla/mux"` 
- installs the dependency we need

`go run <file name>` compiles and runs the code