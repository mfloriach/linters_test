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

func (r commentsMysqlRepository) GetComments(ctx context.Context) (comments []shared.CommentEntity, err error) {
	err = r.db.WithContext(ctx).Find(&comments).Error

	return comments, err
}
