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
	AggregatorHandler *AggregatorHandler
}

func NewHandlersContainer(services *services.ServicesContainer, logger *logger.Logger) *HandlersContainer {
	usersHandler := NewUserHandler(services.UsersSvc, logger)
	productsHandler := NewProductHandler(services.ProductSvc, logger)
	todoHandler := NewTodoHandler(services.TodoSvc, logger)
	commentsHandler := NewCommentsHandler(services.CommentSvc, logger)
	aggregatorHandler := NewAggregatorHandler(services.UsersSvc, services.CommentSvc, services.AggregatorService, logger)
	return &HandlersContainer{
		UsersHandler:      usersHandler,
		ProductsHandler:   productsHandler,
		TodoHandler:       todoHandler,
		CommentsHandler:   commentsHandler,
		AggregatorHandler: aggregatorHandler,
	}
}
