package ports

import (
	config "api-gateway/configs"
	"api-gateway/internal/routes"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type HttpQuoteRepository struct {
	logger  *logger.Logger
	client  *http.Client
	baseUrl string
}

func NewHttpQuoteRepository(httpClient *http.Client, logger *logger.Logger) QuoteRepository {
	return &HttpQuoteRepository{
		logger:  logger,
		client:  httpClient,
		baseUrl: config.Load().BaseURL + routes.PathRoutes.Quotes,
	}
}

func (repo *HttpQuoteRepository) GetQuotes() ([]types.Quote, error) {
	resp, err := repo.client.Get(repo.baseUrl)
	if err != nil {
		return []types.Quote{}, err
	}
	defer resp.Body.Close()

	var result []types.Quote

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []types.Quote{}, err
	}

	return result, nil
}
