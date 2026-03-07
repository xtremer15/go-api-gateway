package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	common_utils "api-gateway/pkg/common-utils"
	"api-gateway/pkg/logger"
	"fmt"
	"strings"
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

	var usersChan <-chan any
	var commmentsChan <-chan any

	if strings.Contains(queryParams, "users") {
		usersChan = svc.UserSvc.FetchAsyncUsers()
		fmt.Println("query paras has users")
	}

	if strings.Contains(queryParams, "comments") {
		commmentsChan = svc.CommsSvc.FetchAsyncComments()
		fmt.Println("query paras has comments")
	}

	if usersChan != nil {
		dataToAggregate["users"] = <-usersChan
	}

	if commmentsChan != nil {
		dataToAggregate["comments"] = <-commmentsChan
	}

	rawUsers := dataToAggregate["users"]
	rawComments := dataToAggregate["comments"]
	typedUsers, okUsers := rawUsers.([]types.User)
	typedComments, okComments := rawComments.([]types.Comment)

	if !okUsers {
		svc.Logger.Error("Not a users slice")
	}

	if !okComments {
		svc.Logger.Error("Not a comments slice")
	}

	for i := range typedUsers {
		typedUsers[i].Comment = typedComments[i]
	}

	delete(dataToAggregate, "comments")

	common_utils.WriteToFileBase64("assets/schemas.json", data)
	return dataToAggregate
}
