package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpProductRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewProductRepo(httpClient *http.Client) ProductsRepo {
	return &HttpProductRepository{
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Products,
	}
}

func (repo *HttpProductRepository) GetProducts() ([]types.Product, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.Product{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Products []types.Product
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.Product{}, err
	}

	return result.Products, nil
}
