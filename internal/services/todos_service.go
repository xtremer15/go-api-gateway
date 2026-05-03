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

func NewToDoService(toDoRepos ports.TodoRepo, logger *logger.Logger) *ToDoService {
	return &ToDoService{
		ToDoRepo: toDoRepos,
		Logger:   logger,
	}
}

func (svc *ToDoService) GetTodos() ([]types.Todo, error) {
	todos, err := svc.ToDoRepo.GetTodos()

	if err != nil {
		svc.Logger.Error("Todos not found")
		return nil, err
	}

	return todos, nil
}

func (svc *ToDoService) FetchAsyncTodos() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)

		todos, err := svc.GetTodos()

		if err != nil {
			svc.Logger.Error(err.Error())
			return
		}
		ch <- todos
	}()

	return ch
}
