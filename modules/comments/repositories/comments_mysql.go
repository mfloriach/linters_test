package repositories

import (
	"context"
	"hoge/modules/comments/shared"

	"gorm.io/gorm"
)

type commentsMysqlRepository struct {
	db gorm.DB
}

func NewCommentsMysqlRepository(db gorm.DB) shared.CommentsRepositoryInterface {
	return commentsMysqlRepository{db: db}
}

func (r commentsMysqlRepository) GetComments(ctx context.Context) string {
	r.db.WithContext(ctx).Exec("SELECT * FROM comments")
	return "posts from mysql db"
}
