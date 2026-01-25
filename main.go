package main

import (
	"api-gateway/internal/handlers"
	"api-gateway/internal/ports"
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// common_utils.ReadFile("assets/schemas.json")
	// fmt.Println("\n")
	// common_utils.ReadJsonFile("envs/dev_env.json")

	log := logger.New("INFO")
	httpClient := &http.Client{}
	var port int = 5200

	// 2. Initialize repository layer
	usersRepository := ports.NewUserHttpRepo(httpClient)
	productsRepository := ports.NewProductRepo(httpClient)
	todosRepository := ports.NewHttpTodoRepo(httpClient)
	commentsRepository := ports.NewHttpCommentRepo(httpClient)

	//3.Init Services layer
	usrSvc := services.NewUserService(usersRepository, log)
	productsSvc := services.NewProductsService(productsRepository, log)
	todosSvc := services.NewToDoService(&todosRepository, log)
	commentsSvc := services.NewCommentsService(&commentsRepository, log)

	//4.Init Handlers layer
	usersHandler := handlers.NewUserHandler(usrSvc, log)
	productsHandler := handlers.NewProductHander(productsSvc, log)
	todoHandler := handlers.NewTodoHandler(todosSvc, log)
	commentsHandler := handlers.NewCommentsHandler(commentsSvc, log)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/users", usersHandler.GetUsers)
	router.Get("/products", productsHandler.GetProducts)
	router.Get("/todos", todoHandler.GetTodos)
	router.Get("/comments", commentsHandler.GetComments)

	fmt.Println("API Gateway is running on port: ", port)
	http.ListenAndServe(":5200", router)

}

func addUserHandler(respWriter http.ResponseWriter, request *http.Request) {
	log := logger.New("INFO")
	rawData, err := io.ReadAll(request.Body)
	parsedJson := bytes.NewBuffer(rawData)

	log.Info("addUserHandler called with body:", map[string]interface{}{
		"body": parsedJson.String(),
	})

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://dummyjson.com/users/add", "application/json", parsedJson)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Read and print the response
	json, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("got data", rawData)
	respWriter.Write(json)
}
