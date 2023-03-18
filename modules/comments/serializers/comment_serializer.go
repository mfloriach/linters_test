package serializers

import (
	"fmt"
	"hoge/modules/comments/shared"
	"hoge/pkg"
)

type comment struct {
	ID      string
	Title   string
	Content string
}

func CommentsSerialize(comments []shared.PostEntity) []comment {
	fmt.Println(len(comments))
	res := make([]comment, len(comments))

	for i, c := range comments {
		res[i] = CommentSerialize(c)
	}

	return res
}

func CommentSerialize(c shared.PostEntity) comment {
	return comment{
		ID:      pkg.HashId(c.ID),
		Title:   c.Title,
		Content: c.Content,
	}
}
