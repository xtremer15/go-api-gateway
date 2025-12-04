package types

type Users struct {
	ID        int    `json:"id validate:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"name" validate:"required,min=1,max=100"`
}
