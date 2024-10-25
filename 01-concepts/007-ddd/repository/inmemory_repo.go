package repository

import (
	"go-tarum/entity"
)

type InmemoryRepo struct {
	posts []entity.Post
}

func NewInmemoryRepository() *InmemoryRepo {
	return &InmemoryRepo{
		posts: []entity.Post{},
	}
}

func (r *InmemoryRepo) Save(post *entity.Post) (*entity.Post, error) {
	r.posts = append(r.posts, *post)
	return post, nil
}

func (r *InmemoryRepo) FindAll() ([]entity.Post, error) {
	return r.posts, nil
}
