package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type RecipeHandler struct {
	Svc    *services.RecipeService
	Logger *logger.Logger
}

func NewRecipeHandler(svc *services.RecipeService, logger *logger.Logger) *RecipeHandler {
	return &RecipeHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (h *RecipeHandler) GetRecipes(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	recipes, err := h.Svc.GetRecipes()

	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(response, "failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(recipes)

	if err != nil {
		h.Logger.Error(err.Error())
	}
}
