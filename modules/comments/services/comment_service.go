package services

import (
	"context"
	"hoge/modules/comments/shared"
)

type commentsService struct {
	commentsRepo shared.CommentsRepositoryInterface
}

func NewCommentsService(commentsRepo shared.CommentsRepositoryInterface) shared.CommentsServiceInterface {
	if commentsRepo == nil {
		panic("commentRepo can not be nil")
	}

	return commentsService{commentsRepo: commentsRepo}
}

func (u commentsService) GetComments(ctx context.Context) (comments []shared.CommentEntity, err error) {
	return u.commentsRepo.GetComments(ctx)
}
