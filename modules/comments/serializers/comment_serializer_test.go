package serializers_test

import (
	"hoge/modules/comments/serializers"
	"hoge/modules/comments/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentSerializer(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		res := serializers.CommentSerialize(shared.CommentEntity{
			ID:      1,
			PostID:  1,
			UserID:  1,
			Content: "my content",
		})

		assert.NotNil(t, res)
	})
}

func TestCommentsSerializer(t *testing.T) {
	t.Run("successfull", func(t *testing.T) {
		res := serializers.CommentsSerialize([]shared.CommentEntity{
			{
				ID:      1,
				PostID:  1,
				UserID:  1,
				Content: "my content",
			},
		})

		assert.NotNil(t, res)
	})
}
