package ports

import (
	"api-gateway/internal/types/types"
)

type UserRepo interface {
	GetUsers() ([]types.User, error)
}

type PostsRepo interface {
	GetPosts() ([]types.Post, error)
}

type CommentRepo interface {
	GetComments() ([]types.Comment, error)
}

type AggregatorRepo interface {
	GetAggregatedData() (any, error)
}

type CartsRepo interface {
	GetCarts() ([]types.Cart, error)
}

type ProductsRepo interface {
	GetProducts() ([]types.Product, error)
}

type TodoRepo interface {
	GetTodos() ([]types.Todo, error)
}

type QuoteRepository interface {
	GetQuotes() ([]types.Quote, error)
}

type RecipeRepository interface {
	GetRecipes() ([]types.Recipe, error)
}
