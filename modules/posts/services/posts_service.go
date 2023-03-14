package services

import (
	"context"
	"hoge/modules/posts/shared"
)

type PostServiceInterface interface {
	GetPosts(ctx context.Context) string
}

type postService struct {
	postRepo shared.PostRepositoryInterface
}

func NewPostService(postRepo shared.PostRepositoryInterface) PostServiceInterface {
	if postRepo == nil {
		panic("sdfdsfdsfds")
	}

	return &postService{postRepo: postRepo}
}

func (s postService) GetPosts(ctx context.Context) string {
	return s.postRepo.GetPosts(ctx)
}
