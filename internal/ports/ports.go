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
