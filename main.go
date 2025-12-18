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

	userSvc := &services.UserService{
		UserRepo: usersRepository,
		Logger:   log,
	}

	productsSvc := &services.ProductsService{
		ProductsRepo: productsRepository,
		Logger:       log,
	}

	productsHandler := &handlers.ProductsHandler{
		Svc:    productsSvc,
		Logger: log,
	}

	handler := &handlers.UsersHandlers{
		Svc:    userSvc,
		Logger: log,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/users", handler.GetUsersHandler)
	router.Get("/products", productsHandler.GetProducts)

	log.Info("Starting server on :5200", nil)
	http.ListenAndServe(":5200", router)

}

func getUsersHandler(respWriter http.ResponseWriter, request *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://dummyjson.com/users")
	if err != nil {
		http.Error(respWriter, "failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(respWriter, "failed to read response", http.StatusInternalServerError)
		return
	}

	respWriter.WriteHeader(resp.StatusCode)
	// fmt.Println(string(data))
	respWriter.Write(data)
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
