package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpTodoRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpTodoRepo(hpttClient *http.Client, logger *logger.Logger) TodoRepo {
	return &HttpTodoRepository{
		logger:  logger,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Todos,
		client:  hpttClient,
	}
}

func (repo *HttpTodoRepository) GetTodos() ([]types.Todo, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.Todo{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Todos []types.Todo `json:"todos"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.Todo{}, err
	}

	return result.Todos, nil
}
