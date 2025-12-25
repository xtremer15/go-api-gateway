package types

type Comment struct {
	ID       int     `json:"id"`
	Body     string  `json:"body"`
	PostID   int     `json:"postId"`
	Likes    int     `json:"likes"`
	Products Product `json:"products"`
}
