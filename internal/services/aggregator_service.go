package services

import (
	registry "api-gateway/internal/async-fetch-registers"
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	merger "api-gateway/internal/utils/merger_engine"
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
	svc.Logger.Info("Received ", map[string]interface{}{
		"data":        data,
		"queryParams": queryParams,
	})
	svc.Logger.Info("Getting aggregated data from multiple services...")

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
		fmt.Println("resourceKey", resourceKey)
		dataToAggregate[resourceKey] = <-ch
	}
	//Will need to add logic of deleting other props related to child in order for the returned
	//objected not to contain the other properties at the whole root level
	rawUsers := dataToAggregate["users"]
	rawComments := dataToAggregate["comments"]
	rawPosts := dataToAggregate["posts"]
	delete(dataToAggregate, "comments")
	delete(dataToAggregate, "posts")

	users, okUsers := rawUsers.([]types.User)
	comments, okComments := rawComments.([]types.Comment)
	posts, okPosts := rawPosts.([]types.Post)

	if okUsers && okComments && okPosts {
		var mergeableUsers []types.MergeableParent
		var mergeableComments []types.MergeableChild
		var mergeablePostsAsParent []types.MergeableParent
		var mergeablePostsAsChild []types.MergeableChild

		for i := range users {
			mergeableUsers = append(mergeableUsers, &users[i])
		}

		for i := range comments {
			mergeableComments = append(mergeableComments, &comments[i])
		}

		merger.Merger(mergeablePostsAsParent, mergeableComments)

		for i := range posts {
			mergeablePostsAsChild = append(mergeablePostsAsChild, &posts[i])
		}

		merger.Merger(mergeableUsers, mergeablePostsAsChild)
	}

	// common_utils.WriteToFileBase64("assets/schemas.json", data)
	return dataToAggregate
}
