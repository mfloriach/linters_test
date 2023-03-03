package repositories

import (
	"context"

	"gorm.io/gorm"
)

type commentsMysqlRepository struct {
	db gorm.DB
}

func NewCommentsMysqlRepository(db gorm.DB) CommentsRepositoryInterface {
	return commentsMysqlRepository{db: db}
}

func (r commentsMysqlRepository) GetComments(ctx context.Context) string {
	return "posts from mysql db"
}
