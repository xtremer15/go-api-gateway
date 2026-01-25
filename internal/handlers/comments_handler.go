package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type CommentHandler struct {
	Svc    *services.CommentService
	Logger *logger.Logger
}

func NewCommentsHandler(svc *services.CommentService, logger *logger.Logger) *CommentHandler {
	return &CommentHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (handler *CommentHandler) GetComments(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	comments, err := handler.Svc.GetComments()

	if err != nil {
		handler.Logger.Error("Failed to fetch comments", map[string]interface{}{
			"error": err.Error(),
		})
		http.Error(response, "Failed to fetch comments", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comments)
}
