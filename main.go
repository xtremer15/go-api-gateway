package main

import (
	"api-gateway/internal/handlers"
	"api-gateway/internal/ports"
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// var dataToWrite = map[string]interface{}{
	// 	"name":  "John Doe",
	// 	"age":   30,
	// 	"email": "test_mail@example.com",
	// }
	// common_utils.ReadFile("assets/schemas.json")
	// common_utils.WriteToFile("assets/schemas.json", dataToWrite)
	// fmt.Println("before===========\n====================")
	// common_utils.ReadFile("assets/schemas.json")
	// common_utils.ReadJsonFile("envs/dev_env.json")

	log := logger.New("INFO")
	httpClient := &http.Client{}
	var port int = 5200

	// 2. Initialize repository layer
	usersRepository := ports.NewUserHttpRepo(httpClient)
	productsRepository := ports.NewProductRepo(httpClient)
	todosRepository := ports.NewHttpTodoRepo(httpClient, log)
	commentsRepository := ports.NewHttpCommentRepo(httpClient, log)
	aggregatorRepository := ports.NewAggregatorRepo(httpClient, log)

	//3.Init Services layer
	usrSvc := services.NewUserService(usersRepository, log)
	productsSvc := services.NewProductsService(productsRepository, log)
	todosSvc := services.NewToDoService(todosRepository, log)
	commentsSvc := services.NewCommentsService(commentsRepository, log)
	aggregatedSvc := services.NewAggregatorService(commentsSvc, usrSvc, aggregatorRepository, log)

	//4.Init Handlers layer
	usersHandler := handlers.NewUserHandler(usrSvc, log)
	productsHandler := handlers.NewProductHander(productsSvc, log)
	todoHandler := handlers.NewTodoHandler(todosSvc, log)
	commentsHandler := handlers.NewCommentsHandler(commentsSvc, log)
	aggregateHandler := handlers.NewAggregatorHandler(usrSvc, commentsSvc, aggregatedSvc, log)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/users", usersHandler.GetUsers)
	router.Get("/products", productsHandler.GetProducts)
	router.Get("/todos", todoHandler.GetTodos)
	router.Get("/comments", commentsHandler.GetComments)
	router.Post("/aggregate", aggregateHandler.AggregateData)

	fmt.Println("API Gateway is running on  nr port: ", port)
	http.ListenAndServe(":5200", router)

}
