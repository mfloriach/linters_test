package serializers

import (
	"hoge/modules/comments/shared"
	"hoge/pkg"
)

type comment struct {
	ID      string
	PostID  string
	UserID  string
	Content string
}

func CommentsSerialize(comments []shared.CommentEntity) []comment {
	res := make([]comment, len(comments))

	for i, c := range comments {
		res[i] = CommentSerialize(c)
	}

	return res
}

func CommentSerialize(c shared.CommentEntity) comment {
	return comment{
		ID:      pkg.HashId(c.ID),
		PostID:  pkg.HashId(c.PostID),
		UserID:  pkg.HashId(c.UserID),
		Content: c.Content,
	}
}
