package shared

import "context"

type CommentsServiceInterface interface {
	GetComments(ctx context.Context) string
}

type CommentsRepositoryInterface interface {
	GetComments(ctx context.Context) string
}
