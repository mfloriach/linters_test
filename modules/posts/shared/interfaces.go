package shared

import "context"

type PostServiceInterface interface {
	GetPosts(ctx context.Context) string
}

type PostRepositoryInterface interface {
	GetPosts(ctx context.Context) string
}
