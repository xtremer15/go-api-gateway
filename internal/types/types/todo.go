package types

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

func (todo *Todo) GetParentID() int {
	return todo.UserID
}

func (todo *Todo) GetChildKey() int {
	return todo.ID
}

func (todo *Todo) GetChildType() string {
	return "todos"
}
