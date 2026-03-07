package handlers

import (
	"api-gateway/internal/services"
	"api-gateway/pkg/logger"
	"encoding/json"
	"net/http"
)

type AggregatorHandler struct {
	UserSvc    *services.UserService
	CommentSvc *services.CommentService
	Svc        *services.AggregatorService
	Logger     *logger.Logger
}

func NewAggregatorHandler(userSvc *services.UserService, commentSvc *services.CommentService, svc *services.AggregatorService, logger *logger.Logger) *AggregatorHandler {
	return &AggregatorHandler{
		UserSvc:    userSvc,
		CommentSvc: commentSvc,
		Svc:        svc,
		Logger:     logger,
	}
}

func (handler *AggregatorHandler) AggregateData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	params := request.URL.Query()

	var payloadData map[string]interface{}
	err := json.NewDecoder(request.Body).Decode(&payloadData)
	queryParams := params.Get("include")
	defer request.Body.Close()

	if err != nil {
		handler.Logger.Error(err.Error())
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	//Payload is not used, is there just in case we will use it in future calls of the POST method from main.go
	//ATM the logic is done via query params
	aggregatedData := handler.Svc.GetAggregatedData(payloadData, queryParams)
	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(&aggregatedData)

	if err != nil {
		handler.Logger.Error(err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

}
