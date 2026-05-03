package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type TodoHandler struct {
	Svc    *services.ToDoService
	Logger *logger.Logger
}

func NewTodoHandler(svc *services.ToDoService, logger *logger.Logger) *TodoHandler {
	return &TodoHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (handler *TodoHandler) GetTodos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	todos, err := handler.Svc.GetTodos()
	if err != nil {
		handler.Logger.Error("failed to fetch todos", map[string]interface{}{
			"error": err.Error(),
		})
		http.Error(response, "failed to fetch todos", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(todos)

}
