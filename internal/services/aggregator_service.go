package services

import (
	"api-gateway/internal/ports"
	"api-gateway/pkg/logger"
	"fmt"
)

type AggregatorService struct {
	AggregatorRepo ports.AggregatorRepo
	Logger         *logger.Logger
}

func NewAggregatorService(aggregatedRepo ports.AggregatorRepo, logger *logger.Logger) *AggregatorService {
	return &AggregatorService{
		AggregatorRepo: aggregatedRepo,
		Logger:         logger,
	}
}

func (svc *AggregatorService) GetAggregatedData() (any, error) {
	fmt.Println("Getting aggregated data from multiple services...")
	return "Implement me", nil
}
