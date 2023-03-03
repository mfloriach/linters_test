package repositories

import "context"

type PostRepositoryInterface interface {
	GetPosts(ctx context.Context) string
}
