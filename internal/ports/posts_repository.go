package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpPostsRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpPostRepo(httpClient *http.Client, logger *logger.Logger) PostsRepo {
	return &HttpPostsRepository{
		logger:  logger,
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Posts + "?limit=240",
	}
}

func (repo *HttpPostsRepository) GetPosts() ([]types.Post, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.Post{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Posts []types.Post `json:"posts"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.Post{}, err
	}

	return result.Posts, nil
}
