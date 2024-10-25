## Steps

- initialize a mod --> `go mod init <name>`
- check this repo to add the dependency --> github.com/gocolly/colly
- download the dependencies --> `go mod download`
  - this should create a go.sum file with all the subdependencies listed
- run the project by doing `go run .`