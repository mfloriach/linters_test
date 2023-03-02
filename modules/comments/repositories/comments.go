package repositories

import "context"

type CommentsRepository interface {
	GetComments(ctx context.Context) string
}
