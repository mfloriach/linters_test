package shared

type CommentEntity struct {
	ID      uint
	UserID  uint
	PostID  uint
	Content string
}

func (CommentEntity) TableName() string {
	return "comments"
}
