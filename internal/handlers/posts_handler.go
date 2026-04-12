package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type PostsHandler struct {
	Svc    *services.PostsService
	Logger *logger.Logger
}

func NewPostsHandler(svc *services.PostsService, logger *logger.Logger) *PostsHandler {
	return &PostsHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (handler *PostsHandler) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	posts, err := handler.Svc.GetPosts()

	if err != nil {
		handler.Logger.Error("Failed to fetch posts", map[string]interface{}{
			"error": err.Error(),
		})
		http.Error(response, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
	handler.Logger.Info("Posts fetched successfully", map[string]interface{}{
		"posts": posts,
	})
}
