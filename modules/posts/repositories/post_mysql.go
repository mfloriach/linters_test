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

func (r postMysqlRepository) GetPosts(ctx context.Context) (posts []shared.PostEntity, err error) {
	err = r.db.WithContext(ctx).Find(&posts).Error

	return posts, err
}
