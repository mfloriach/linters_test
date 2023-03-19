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

func (s postService) GetPosts(ctx context.Context) ([]shared.PostEntity, error) {
	return s.postRepo.GetPosts(ctx)
}
