package main

import (
	registry "api-gateway/internal/async-fetch-registers"

	"api-gateway/internal/handlers"
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

	//Init Services layer
	services := services.NewServiceContainer(httpClient, log)

	registry.RegisterAsyncFn("users", services.UsersSvc.FetchAsyncUsers)
	registry.RegisterAsyncFn("comments", services.CommentSvc.FetchAsyncComments)
	//4.Init Handlers layer
	handlers := handlers.NewHandlersContainer(services, log)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/users", handlers.UsersHandler.GetUsers)
	router.Get("/products", handlers.ProductsHandler.GetProducts)
	router.Get("/todos", handlers.TodoHandler.GetTodos)
	router.Get("/comments", handlers.CommentsHandler.GetComments)
	router.Post("/aggregate", handlers.AggregatorHandler.AggregateData)

	fmt.Println("API Gateway is running on  nr port: ", port)
	http.ListenAndServe(":5200", router)

}
