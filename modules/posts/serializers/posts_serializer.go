package serializers

import (
	"hoge/modules/posts/shared"
	"hoge/pkg"
)

type post struct {
	ID      string
	Title   string
	Content string
}

func PostsSerialize(posts []shared.PostEntity) []post {
	res := make([]post, len(posts))

	for i, p := range posts {
		res[i] = PostSerialize(p)
	}

	return res
}

func PostSerialize(p shared.PostEntity) post {
	return post{
		ID:      pkg.HashId(p.ID),
		Title:   p.Title,
		Content: p.Content,
	}
}
