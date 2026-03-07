package types

type AggregatorRequest struct {
	Users    User    `json:"users"`
	Products Product `json:"products"`
	Todos    Todo    `json:"todos"`
	Comments Comment `json:"comments"`
}
