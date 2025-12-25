package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type ToDoService struct {
	ToDoRepo ports.TodoRepo
	Logger   *logger.Logger
}

func (svc *ToDoService) GetProducts() ([]types.Todo, error) {
	todos, err := svc.ToDoRepo.GetTodos()

	if err != nil {
		svc.Logger.Error("Todos not found")
		return nil, err
	}

	return todos, nil
}
