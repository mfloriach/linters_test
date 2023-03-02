package repositories

import "context"

type PostRepository interface {
	GetPosts(ctx context.Context) string
}
