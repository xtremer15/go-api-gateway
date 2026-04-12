package ports

import (
	"api-gateway/internal/types/types"
)

type UserRepo interface {
	GetUsers() ([]types.User, error)
}

type ProductsRepo interface {
	GetProducts() ([]types.Product, error)
}

type TodoRepo interface {
	GetTodos() ([]types.Todo, error)
}

type CommentRepo interface {
	GetComments() ([]types.Comment, error)
}

type AggregatorRepo interface {
	GetAggregatedData() (any, error)
}

type PostsRepo interface {
	GetPosts() ([]types.Post, error)
}
