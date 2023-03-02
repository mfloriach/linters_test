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
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostServiceInterface {
	return &postService{postRepo: postRepo}
}

func (s postService) GetPosts(ctx context.Context) string {
	r := gin.Default()
	db := pkg.GormSqliteSetup("gfdgfd")
	comments.GinServer(r, db)
	return s.postRepo.GetPosts(ctx)
}
