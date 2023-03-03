package repositories

import "context"

type CommentsRepositoryInterface interface {
	GetComments(ctx context.Context) string
}
