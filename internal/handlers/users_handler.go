package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type UsersHandlers struct {
	Svc    *services.UserService
	Logger *logger.Logger
}

func (handler *UsersHandlers) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	users, err := handler.Svc.GetUsers()

	if err != nil {
		handler.Logger.Error("failed to fetch users", map[string]interface{}{
			"error": err.Error(),
		})
		http.Error(response, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}
