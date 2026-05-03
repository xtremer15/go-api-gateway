package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
)

type HandlersContainer struct {
	UsersHandler      *UsersHandlers
	ProductsHandler   *ProductsHandler
	TodoHandler       *TodoHandler
	CommentsHandler   *CommentHandler
	PostsHandler      *PostsHandler
	AggregatorHandler *AggregatorHandler
	QuoteHandler      *QuoteHandler
	CartHandler       *CartHandler
	RecipeHandler     *RecipeHandler
}

func NewHandlersContainer(services *services.ServicesContainer, logger *logger.Logger) *HandlersContainer {
	usersHandler := NewUserHandler(services.UsersSvc, logger)
	productsHandler := NewProductHandler(services.ProductSvc, logger)
	todoHandler := NewTodoHandler(services.TodoSvc, logger)
	commentsHandler := NewCommentsHandler(services.CommentSvc, logger)
	postsHandler := NewPostsHandler(services.PostsSvc, logger)
	quoteHandler := NewQuoteHandler(services.QuoteSvc, logger)
	cartHandler := NewCartHandler(services.CartSvc, logger)
	recipeHandler := NewRecipeHandler(services.RecipeSvc, logger)
	aggregatorHandler := NewAggregatorHandler(services.UsersSvc, services.CommentSvc, services.AggregatorService, logger)
	return &HandlersContainer{
		UsersHandler:      usersHandler,
		ProductsHandler:   productsHandler,
		TodoHandler:       todoHandler,
		CommentsHandler:   commentsHandler,
		PostsHandler:      postsHandler,
		QuoteHandler:      quoteHandler,
		CartHandler:       cartHandler,
		RecipeHandler:     recipeHandler,
		AggregatorHandler: aggregatorHandler,
	}
}
