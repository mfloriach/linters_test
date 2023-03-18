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

func (r commentsMysqlRepository) GetComments(ctx context.Context) []shared.PostEntity {
	r.db.WithContext(ctx).Exec("SELECT * FROM comments")
	return []shared.PostEntity{
		{
			ID:      1,
			Title:   "post 1",
			Content: "content 1",
		},
		{
			ID:      2,
			Title:   "post 2",
			Content: "content 2",
		},
	}
}
