## Commands
`go mod init go-3-bookstore`
- initialized a go.mod file
- Describes the module: with name/module path go version used in the program
- The module path is also the import path (e.g. import "github/nana/booking-app")
- Creates a new module
- Module path can correspond to a repository you plan to publish your module.

`go get "github.com/jinzhu/gorm"` --> let us interact with the DB, it is a ORM
`go get "github.com/jinzhu/gorm/dialects/mysql"` --> the DB
`go get "github.com/gorilla/mux"` --> controls our routing
- installs the dependency we need

`go run <file name>` compiles and runs the code