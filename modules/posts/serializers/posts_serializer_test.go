package serializers_test

import (
	"hoge/modules/posts/serializers"
	"hoge/modules/posts/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentSerializer(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		res := serializers.PostSerialize(shared.PostEntity{
			ID:      1,
			Title:   "title",
			Content: "my content",
		})

		assert.NotNil(t, res)
	})
}

func TestCommentsSerializer(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		res := serializers.PostsSerialize([]shared.PostEntity{
			{
				ID:      1,
				Title:   "title",
				Content: "my content",
			},
		})

		assert.NotNil(t, res)
	})
}
