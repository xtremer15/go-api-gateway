package services

import (
	registry "api-gateway/internal/async-fetch-registers"
	"api-gateway/internal/ports"
	common_utils "api-gateway/pkg/common-utils"
	"api-gateway/pkg/logger"
	"fmt"
)

type AggregatorService struct {
	AggregatorRepo ports.AggregatorRepo
	Logger         *logger.Logger
	UserSvc        *UserService
	CommsSvc       *CommentService
}

func NewAggregatorService(commsSvc *CommentService, userSvc *UserService, aggregatedRepo ports.AggregatorRepo, logger *logger.Logger) *AggregatorService {
	return &AggregatorService{
		AggregatorRepo: aggregatedRepo,
		Logger:         logger,
		UserSvc:        userSvc,
		CommsSvc:       commsSvc,
	}
}

func (svc *AggregatorService) GetAggregatedData(data map[string]interface{}, queryParams string) any {
	fmt.Println("recived", data, queryParams)
	fmt.Println("Getting aggregated data from multiple services...")

	var dataToAggregate map[string]any = make(map[string]any)
	activeChannels := make(map[string]<-chan any)

	for key := range registry.GlobalRegistry {
		fn := registry.GetRegisteredFn(key)
		if fn != nil {
			// Execute it dynamically and save the channel!
			activeChannels[key] = fn()
		} else {
			svc.Logger.Error("Client requested unknown resource: " + key)
		}

	}

	for resourceKey, ch := range activeChannels {
		dataToAggregate[resourceKey] = <-ch
	}
	delete(dataToAggregate, "comments")

	common_utils.WriteToFileBase64("assets/schemas.json", data)
	return dataToAggregate
}
