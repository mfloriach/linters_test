package services

import (
	"context"
	"hoge/modules/comments"
	"hoge/modules/posts/repositories"
	"hoge/pkg"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	db := pkg.GormSqliteSetup("gfdgfd")
	comments.GinServer(r, db)
	return s.postRepo.GetPosts(ctx)
}
