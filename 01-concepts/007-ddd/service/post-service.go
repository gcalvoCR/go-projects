package service

import (
	"errors"
	"go-tarum/entity"
	"go-tarum/repository"
	"math/rand"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

var (
	repo repository.PostRepository
)

type service struct{}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}
	if post.Title == "" {
		return errors.New("title is empty")
	}
	if post.Text == "" {
		return errors.New("text is empty")
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Intn(10000000)
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
