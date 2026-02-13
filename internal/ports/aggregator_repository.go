package ports

import (
	config "api-gateway/configs"
	"api-gateway/pkg/logger"
	"fmt"
	"net/http"
)

type AggregatorRepository struct {
	httpClient *http.Client
	logger     *logger.Logger
	baseUrl    string
}

func NewAggregatorRepo(httpClient *http.Client, logger *logger.Logger) AggregatorRepo {
	return &AggregatorRepository{
		httpClient: httpClient,
		logger:     logger,
		baseUrl:    config.Load().BaseURL + "/aggregate",
	}
}

// GetAggregatedData implements [AggregatroRepo].
func (a *AggregatorRepository) GetAggregatedData() (any, error) {
	fmt.Println("Getting aggregated data from multiple services...")
	return "Implement me", nil
}
