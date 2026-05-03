package types

type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
	UserID int    `json:"userId"`
}

func (q *Quote) GetParentID() int {
	return q.UserID
}
func (q *Quote) GetChildKey() int {
	return q.ID
}
func (q *Quote) GetChildType() string {
	return "quote"
}
