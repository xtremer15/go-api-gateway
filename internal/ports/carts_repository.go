package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpCartsRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpCartsRepo(httpClient *http.Client, logger *logger.Logger) CartsRepo {
	return &HttpCartsRepository{
		logger:  logger,
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Carts,
	}

}

func (repo *HttpCartsRepository) GetCarts() ([]types.Cart, error) {
	carts, err := repo.client.Get(repo.baseUrl)

	if err != nil {
		return []types.Cart{}, err
	}

	defer carts.Body.Close()

	var result []types.Cart

	err = json.NewDecoder(carts.Body).Decode(&result)

	if err != nil {
		return []types.Cart{}, err
	}

	return result, nil
}
