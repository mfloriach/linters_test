package shared

type PostEntity struct {
	ID      uint
	Title   string
	Content string
}

func (PostEntity) TableName() string {
	return "posts"
}
