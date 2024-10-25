package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct{}

var chiDispatcher = chi.NewRouter()

func NewChiRouter() Router {
	return &chiRouter{}
}

func (c *chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (c *chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (c *chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP Server running on port", port)
	http.ListenAndServe(port, chiDispatcher)
}
