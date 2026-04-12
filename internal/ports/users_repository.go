package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpUsersRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewUserHttpRepo(httpClient *http.Client) UserRepo {
	return &HttpUsersRepository{ // -> *HttpUsersRepository
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Users + "?limit=240",
		client:  httpClient,
	}
}

func (repo *HttpUsersRepository) GetUsers() ([]types.User, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.User{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Users []types.User `json:"users"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.User{}, err
	}

	return result.Users, nil

}
