package shared

import "context"

type CommentsServiceInterface interface {
	GetComments(ctx context.Context) []PostEntity
}

type CommentsRepositoryInterface interface {
	GetComments(ctx context.Context) []PostEntity
}
