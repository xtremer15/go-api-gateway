package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpRecipeRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpRecipeRepository(httpClient *http.Client, logger *logger.Logger) RecipeRepository {
	return &HttpRecipeRepository{
		logger:  logger,
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Recipes,
	}
}

func (repo *HttpRecipeRepository) GetRecipes() ([]types.Recipe, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.Recipe{}, err
	}
	defer resp.Body.Close()

	var result []types.Recipe

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.Recipe{}, err
	}

	return result, nil
}
