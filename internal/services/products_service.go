package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type ProductsService struct {
	ProductsRepo ports.ProductsRepo
	Logger       *logger.Logger
}

func NewProductsService(productsRepo ports.ProductsRepo, logger *logger.Logger) *ProductsService {
	return &ProductsService{
		ProductsRepo: productsRepo,
		Logger:       logger,
	}
}

func (svc *ProductsService) GetProducts() ([]types.Product, error) {
	products, err := svc.ProductsRepo.GetProducts()

	if err != nil {
		svc.Logger.Error("Products not found")
		return nil, err
	}

	return products, nil
}
