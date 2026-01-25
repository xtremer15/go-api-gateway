package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type ProductsHandler struct {
	Svc    *services.ProductsService
	Logger *logger.Logger
}

func NewProductHander(svc *services.ProductsService, logger *logger.Logger) *ProductsHandler {
	return &ProductsHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (handler *ProductsHandler) GetProducts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	products, err := handler.Svc.GetProducts()

	if err != nil {
		handler.Logger.Error("failed to fetch users", map[string]interface{}{
			"error": err.Error(),
		})
		http.Error(response, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(products)
}
