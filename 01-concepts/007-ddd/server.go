package main

import (
	"go-tarum/controller"
	router "go-tarum/http"
	"go-tarum/repository"
	"go-tarum/service"
	"net/http"
)

var (
	repo           repository.PostRepository = repository.NewInmemoryRepository()
	postService    service.PostService       = service.NewPostService(repo)
	httpRouter     router.Router             = router.NewChiRouter()
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	port := ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
