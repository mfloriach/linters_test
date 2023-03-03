package serializers

import "time"

type comment struct {
	ID        string
	UserID    string
	PostID    string
	Text      string
	CreatedAt time.Time
}

func CommentsSerialize(comments string) []comment {
	return []comment{
		{
			ID:        "asdsadas",
			UserID:    "dfdsfdsf",
			PostID:    "dsfdsfdsfsdf",
			Text:      "dsfdsfdsf",
			CreatedAt: time.Now(),
		},
	}
}
