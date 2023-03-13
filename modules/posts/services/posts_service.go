package services

import (
	"context"
	"hoge/modules/posts/repositories"
)

type PostServiceInterface interface {
	GetPosts(ctx context.Context) string
}

type postService struct {
	postRepo repositories.PostRepositoryInterface
}

func NewPostService(postRepo repositories.PostRepositoryInterface) PostServiceInterface {
	if postRepo == nil {
		panic("sdfdsfdsfds")
	}

	return &postService{postRepo: postRepo}
}

func (s postService) GetPosts(ctx context.Context) string {
	return s.postRepo.GetPosts(ctx)
}
