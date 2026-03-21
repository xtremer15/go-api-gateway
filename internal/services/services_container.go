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
	AggregatorService *AggregatorService
}

func NewServiceContainer(httpClient *http.Client, logger *logger.Logger) *ServicesContainer {
	// Instantiate Repos
	usersRepository := ports.NewUserHttpRepo(httpClient)
	productsRepository := ports.NewProductRepo(httpClient)
	todosRepository := ports.NewHttpTodoRepo(httpClient, logger)
	commentsRepository := ports.NewHttpCommentRepo(httpClient, logger)
	aggregatorRepository := ports.NewAggregatorRepo(httpClient, logger)

	// Instantiate Services
	userSvc := NewUserService(usersRepository, logger)
	productSvc := NewProductsService(productsRepository, logger)
	todoSvc := NewToDoService(todosRepository, logger)
	commentSvc := NewCommentsService(commentsRepository, logger)
	aggregatorSvc := NewAggregatorService(commentSvc, userSvc, aggregatorRepository, logger)
	// Return the container!
	return &ServicesContainer{
		UsersSvc:          userSvc,
		ProductSvc:        productSvc,
		TodoSvc:           todoSvc,
		CommentSvc:        commentSvc,
		AggregatorService: aggregatorSvc,
	}
}
