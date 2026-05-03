package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	CartSvc *services.CartsService
	Logger  *logger.Logger
}

func NewCartHandler(cartSvc *services.CartsService, logger *logger.Logger) *CartHandler {
	return &CartHandler{
		CartSvc: cartSvc,
		Logger:  logger,
	}
}

func (h *CartHandler) GetCarts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	carts, err := h.CartSvc.GetCarts()

	if carts != nil {
		h.Logger.Error(err.Error())
	}

	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(carts)

	if err != nil {
		h.Logger.Error(err.Error())
	}
}
