package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpCommentsRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpCommentRepo(httpClient *http.Client) CommentRepo {
	return &HttpCommentsRepository{
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Comments,
	}
}

func (repo *HttpCommentsRepository) GetComments() ([]types.Comment, error) {
	resp, err := repo.client.Get(repo.baseUrl)

	if err != nil {
		return []types.Comment{}, err
	}

	defer resp.Body.Close()

	var result struct {
		Comments []types.Comment `json:"comments"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return []types.Comment{}, err
	}

	return result.Comments, nil
}
