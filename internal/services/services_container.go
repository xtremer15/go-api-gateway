package services

import (
	"api-gateway/internal/ports"
	"api-gateway/pkg/logger"
	"net/http"
)

type ServicesContainer struct {
	UsersSvc          *UserService
	CommentSvc        *CommentService
	TodoSvc           *ToDoService
	ProductSvc        *ProductsService
	PostsSvc          *PostsService
	AggregatorService *AggregatorService
	CartSvc           *CartsService
	QuoteSvc          *QuoteService
	RecipeSvc         *RecipeService
}

func NewServiceContainer(httpClient *http.Client, logger *logger.Logger) *ServicesContainer {
	// Instantiate Repos
	usersRepository := ports.NewUserHttpRepo(httpClient)
	productsRepository := ports.NewProductRepo(httpClient)
	todosRepository := ports.NewHttpTodoRepo(httpClient, logger)
	commentsRepository := ports.NewHttpCommentRepo(httpClient, logger)
	postsRepository := ports.NewHttpPostRepo(httpClient, logger)
	aggregatorRepository := ports.NewAggregatorRepo(httpClient, logger)
	cartsRepository := ports.NewHttpCartsRepo(httpClient, logger)
	quotesRepository := ports.NewHttpQuoteRepository(httpClient, logger)
	recipesRepository := ports.NewHttpRecipeRepository(httpClient, logger)
	// Instantiate Services
	userSvc := NewUserService(usersRepository, logger)
	productSvc := NewProductsService(productsRepository, logger)
	todoSvc := NewToDoService(todosRepository, logger)
	commentSvc := NewCommentsService(commentsRepository, logger)
	postsSvc := NewPostsService(postsRepository, logger)
	quoteSvc := NewQuoteService(quotesRepository, logger)
	cartSvc := NewCartsService(cartsRepository, logger)
	recipeSvc := NewRecipeService(recipesRepository, logger)
	aggregatorSvc := NewAggregatorService(commentSvc, userSvc, aggregatorRepository, logger)
	// Return the container!
	return &ServicesContainer{
		UsersSvc:          userSvc,
		ProductSvc:        productSvc,
		TodoSvc:           todoSvc,
		CommentSvc:        commentSvc,
		PostsSvc:          postsSvc,
		QuoteSvc:          quoteSvc,
		CartSvc:           cartSvc,
		RecipeSvc:         recipeSvc,
		AggregatorService: aggregatorSvc,
	}
}
