package types

type Comment struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	PostID int    `json:"postId"`
	Likes  int    `json:"likes"`
}

// Interfaces implementation
func (comment *Comment) GetParentID() int {
	// Might need to change it with post ID
	return comment.PostID
}

func (comment *Comment) GetChildKey() int {
	return comment.ID
}

func (comment *Comment) GetChildType() string {
	return "comments"
}
