package services

import (
	"context"
	"hoge/modules/posts/shared"
)

type postService struct {
	postRepo shared.PostRepositoryInterface
}

func NewPostService(postRepo shared.PostRepositoryInterface) shared.PostServiceInterface {
	if postRepo == nil {
		panic("post repository can not be nil")
	}

	return &postService{postRepo: postRepo}
}

func (s postService) GetPosts(ctx context.Context) string {
	return s.postRepo.GetPosts(ctx)
}
