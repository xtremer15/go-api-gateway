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

	// 2. Initialize repository layer
	usersRepository := ports.NewUserHttpRepo(httpClient)
	productsRepository := ports.NewProductRepo(httpClient)
	todosRepository := ports.NewHttpTodoRepo(httpClient)
	commentsRepository := ports.NewHttpCommentRepo(httpClient)

	//3.Init Services layer
	//TODO: Move init of the service in the services layer of each service by creating a function instead of
	//creating a struct with defaults values
	userSvc := &services.UserService{
		UserRepo: usersRepository,
		Logger:   log,
	}

	productsSvc := &services.ProductsService{
		ProductsRepo: productsRepository,
		Logger:       log,
	}

	todosSvc := &services.ToDoService{
		ToDoRepo: todosRepository,
		Logger:   log,
	}

	commentsSvc := &services.CommentService{
		CommentsRepo: commentsRepository,
		Logger:       log,
	}

	//4.Init Handlers layer
	//TODO: Move init of the repo in the repo layer of each repo by creating a function instead of
	//creating a struct with defaults values
	productsHandler := &handlers.ProductsHandler{
		Svc:    productsSvc,
		Logger: log,
	}

	usersHandler := &handlers.UsersHandlers{
		Svc:    userSvc,
		Logger: log,
	}

	todoHandler := &handlers.TodoHandler{
		Svc:    todosSvc,
		Logger: log,
	}

	commentsHandler := &handlers.CommentHandler{
		Svc:    commentsSvc,
		Logger: log,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/users", usersHandler.GetUsers)
	router.Get("/products", productsHandler.GetProducts)
	router.Get("/todos", todoHandler.GetTodos)
	router.Get("/comments", commentsHandler.GetComments)

	log.Info("Starting server on :5200", nil)
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
