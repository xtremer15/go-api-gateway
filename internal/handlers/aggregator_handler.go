package handlers

import (
	"api-gateway/internal/services"
	common_utils "api-gateway/pkg/common-utils"
	"api-gateway/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AggregatorHandler struct {
	Svc    *services.AggregatorService
	Logger *logger.Logger
}

func NewAggregatorHandler(svc *services.AggregatorService, logger *logger.Logger) *AggregatorHandler {
	return &AggregatorHandler{
		Svc:    svc,
		Logger: logger,
	}
}

func (handler *AggregatorHandler) AggregateData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var payloadData map[string]interface{}

	_, err := handler.Svc.GetAggregatedData()

	defer request.Body.Close()
	body, err := io.ReadAll(request.Body)

	fmt.Println("BODY LEN:", len(body))
	fmt.Println("RAW BODY:", string(body))

	if err != nil {
		handler.Logger.Error(err.Error())
	}

	if err != nil {
		handler.Logger.Error(err.Error())
		return
	}

	if err := json.Unmarshal(body, &payloadData); err != nil {
		handler.Logger.Error(err.Error())
		return
	}

	formatted, err := json.MarshalIndent(payloadData, "", "  ")
	if err != nil {
		handler.Logger.Error(err.Error())
		return
	}

	if err := os.WriteFile("assets/schemas.json", formatted, 0644); err != nil {
		handler.Logger.Error(err.Error())
		return
	}

	response.Write(body)
	common_utils.WriteToFileBase64("assets/schemas.json", formatted)
	fmt.Println("Data written to file successfully", string(body))

}
