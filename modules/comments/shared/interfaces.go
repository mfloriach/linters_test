package shared

import "context"

type CommentsServiceInterface interface {
	GetComments(ctx context.Context) (comments []CommentEntity, err error)
}

type CommentsRepositoryInterface interface {
	GetComments(ctx context.Context) (comments []CommentEntity, err error)
}
