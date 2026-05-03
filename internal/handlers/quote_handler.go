package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type QuoteHandler struct {
	Logger *logger.Logger
	Svc    *services.QuoteService
}

func NewQuoteHandler(svc *services.QuoteService, logger *logger.Logger) *QuoteHandler {
	return &QuoteHandler{
		Logger: logger,
		Svc:    svc,
	}
}

func (h *QuoteHandler) GetQuotes(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	quotes, err := h.Svc.GetQuotes()

	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(response, "failed to fetch quotes", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(quotes)

	if err != nil {
		h.Logger.Error(err.Error())
	}
}
