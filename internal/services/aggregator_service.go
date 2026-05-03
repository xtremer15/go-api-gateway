package services

import (
	registry "api-gateway/internal/async-fetch-registers"
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	merger "api-gateway/internal/utils/merger_engine"
	"api-gateway/pkg/logger"
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

func (svc *AggregatorService) GetAggregatedData(data map[string]interface{}, queryParams []string) any {
	svc.Logger.Info("Received ", map[string]interface{}{
		"data":        data,
		"queryParams": queryParams,
	})
	svc.Logger.Info("Getting aggregated data from multiple services...")

	var dataToAggregate map[string]any = make(map[string]any)
	activeChannels := make(map[string]<-chan any)

	for _, key := range queryParams {
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
	//Will need to add logic of deleting other props related to child in order for the returned
	//objected not to contain the other properties at the whole root level
	rawUsers := dataToAggregate["users"]
	rawComments := dataToAggregate["comments"]
	rawPosts := dataToAggregate["posts"]
	rawCarts := dataToAggregate["carts"]
	rawQuotes := dataToAggregate["quotes"]
	rawRecipes := dataToAggregate["recipes"]

	delete(dataToAggregate, "comments")
	delete(dataToAggregate, "posts")
	delete(dataToAggregate, "carts")
	delete(dataToAggregate, "quotes")
	delete(dataToAggregate, "recipes")

	users, okUsers := rawUsers.([]types.User)
	comments, okComments := rawComments.([]types.Comment)
	posts, okPosts := rawPosts.([]types.Post)
	carts, okCarts := rawCarts.([]types.Cart)
	quotes, okQuotes := rawQuotes.([]types.Quote)
	recipes, okRecipes := rawRecipes.([]types.Recipe)

	var mergeableUsers []types.MergeableParent
	var mergeableComments []types.MergeableChild

	var mergeablePostsAsParent []types.MergeableParent
	var mergeablePostsAsChild []types.MergeableChild

	var mergeableCartsAsParent []types.MergeableParent
	var mergeableCartsAsChild []types.MergeableChild

	var mergeableQuotes []types.MergeableChild
	var mergeableRecipes []types.MergeableChild

	if okPosts && okComments {

		for i := range comments {
			mergeableComments = append(mergeableComments, &comments[i])
		}

		for i := range posts {
			mergeablePostsAsParent = append(mergeablePostsAsParent, &posts[i])
		}

		merger.Merger(mergeablePostsAsParent, mergeableComments)

	}

	if okUsers && okPosts {

		for i := range users {
			mergeableUsers = append(mergeableUsers, &users[i])
		}

		for i := range posts {
			mergeablePostsAsChild = append(mergeablePostsAsChild, &posts[i])
		}

		merger.Merger(mergeableUsers, mergeablePostsAsChild)
	}

	if okCarts && okUsers {

		for i := range carts {
			mergeableCarts = append(mergeableCarts, &carts[i])
		}

		for i := range users {
			mergeableUsers = append(mergeableUsers, &users[i])
		}

		merger.Merger(mergeableUsers, mergeableCarts)
	}

	if okQuotes && okUsers {

		for i := range quotes {
			mergeableQuotes = append(mergeableQuotes, &quotes[i])
		}

		for i := range users {
			mergeableUsers = append(mergeableUsers, &users[i])
		}

		merger.Merger(mergeableUsers, mergeableQuotes)
	}

	if okRecipes && okUsers {

		for i := range recipes {
			mergeableRecipes = append(mergeableRecipes, &recipes[i])
		}

		for i := range users {
			mergeableUsers = append(mergeableUsers, &users[i])
		}

		merger.Merger(mergeableUsers, mergeableRecipes)
	}
	// common_utils.WriteToFileBase64("assets/schemas.json", data)
	return dataToAggregate
}
