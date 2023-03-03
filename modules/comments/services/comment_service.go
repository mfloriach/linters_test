package services

import (
	"context"
	"hoge/modules/comments/repositories"
)

type CommentsServiceInterface interface {
	GetComments(ctx context.Context) string
}

type commentsService struct {
	commentsRepo repositories.CommentsRepositoryInterface
}

func NewCommentsService(commentsRepo repositories.CommentsRepositoryInterface) CommentsServiceInterface {
	if commentsRepo == nil {
		panic("sdfdsfdsfds")
	}

	return commentsService{commentsRepo: commentsRepo}
}

func (u commentsService) GetComments(ctx context.Context) string {
	return u.commentsRepo.GetComments(ctx)
}
