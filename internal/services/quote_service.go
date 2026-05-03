package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type QuoteService struct {
	QuoteRepo ports.QuoteRepository
	Logger    *logger.Logger
}

func NewQuoteService(quoteRepo ports.QuoteRepository, logger *logger.Logger) *QuoteService {
	return &QuoteService{
		QuoteRepo: quoteRepo,
		Logger:    logger,
	}
}

func (svc *QuoteService) GetQuotes() ([]types.Quote, error) {
	quotes, err := svc.QuoteRepo.GetQuotes()

	if err != nil {
		svc.Logger.Error("Quotes not found")
		return nil, err
	}

	return quotes, nil
}

func (svc *QuoteService) FetchAsyncQuotes() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)
		quotes, err := svc.GetQuotes()
		if err != nil {
			svc.Logger.Error(err.Error())
			return
		}
		ch <- quotes
	}()

	return ch
}
