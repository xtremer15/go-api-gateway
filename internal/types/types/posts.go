package types

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Tags      []string  `json:"tags"`
	Reactions Reactions `json:"reactions"`
	Views     int       `json:"views"`
	UserID    int       `json:"userId"`
	CommentID int       `json:"commentId"`
	Comments  []Comment `json:"comments"`
}

type Reactions struct {
	Likes    int `json:"likes"`
	Dislikes int `json:"dislikes"`
}

// =============MergeableParent Interface impl. starts=========
// Exposed for parent (User)
func (post *Post) GetResourceID() int {
	return post.ID
}

func (post *Post) AttachChildren(children any, relation string) any {

	switch relation {
	case "comments":
		var concreteComms []Comment
		unpackedComms, _ := children.([]MergeableChild)

		for idx := range unpackedComms {
			concreteComms = append(concreteComms, *unpackedComms[idx].(*Comment))
		}

		post.Comments = concreteComms
	}

	return nil
}

func (post *Post) GetChildren(childrenID int) any {
	return post.Comments[childrenID].GetChildKey()
}

//=============MergeableParent Interface impl. ends=========

// =============MergeableChild Interface impl. starts=========
// Exposed for children (Comments)
func (post *Post) GetParentID() int {
	return post.UserID
}

func (post *Post) GetChildKey() int {
	return post.ID
}

func (post *Post) GetChildType() string {
	return "posts"
}

//=============MergeableChild Interface impl. ends=========
