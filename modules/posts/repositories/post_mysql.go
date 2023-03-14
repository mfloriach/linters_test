package repositories

import (
	"context"
	"hoge/modules/posts/shared"

	"gorm.io/gorm"
)

type postMysqlRepository struct {
	db gorm.DB
}

func NewPostMysqlRepository(db gorm.DB) shared.PostRepositoryInterface {
	return postMysqlRepository{db: db}
}

func (r postMysqlRepository) GetPosts(ctx context.Context) string {
	return "posts from mysql db"
}
