package services

import (
	"context"
	"hoge/modules/comments/repositories"
)

type CommentsServiceInterface interface {
	GetComments(ctx context.Context) string
}

type commentsService struct {
	commentsRepo repositories.CommentsRepository
}

func NewCommentsService(commentsRepo repositories.CommentsRepository) CommentsServiceInterface {
	return commentsService{commentsRepo: commentsRepo}
}

func (u commentsService) GetComments(ctx context.Context) string {
	return u.commentsRepo.GetComments(ctx)
}
