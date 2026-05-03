package types

type AggregatorRequest struct {
	Users    User    `json:"users"`
	Products Product `json:"products"`
	Todos    Todo    `json:"todos"`
	Comments Comment `json:"comments"`
	Posts    Post    `json:"posts"`
	Carts    Cart    `json:"carts"`
	Quotes   Quote   `json:"quotes"`
	Recipes  Recipe  `json:"recipes"`
}
