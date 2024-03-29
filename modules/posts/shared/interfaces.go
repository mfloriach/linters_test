package shared

import (
	"context"
)

type PostServiceInterface interface {
	GetPosts(ctx context.Context) ([]PostEntity, error)
}

type PostRepositoryInterface interface {
	GetPosts(ctx context.Context) ([]PostEntity, error)
}
