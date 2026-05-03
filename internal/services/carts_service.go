package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type CartsService struct {
	CartsRepo ports.CartsRepo
	Logger    *logger.Logger
}

func NewCartsService(cartsRepo ports.CartsRepo, logger *logger.Logger) *CartsService {
	return &CartsService{
		CartsRepo: cartsRepo,
		Logger:    logger,
	}
}

func (svc *CartsService) GetCarts() ([]types.Cart, error) {
	carts, err := svc.CartsRepo.GetCarts()

	if err != nil {
		svc.Logger.Error("Carts not found")
		return nil, err
	}

	return carts, nil
}

func (svc *CartsService) FetchAsyncCarts() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)
		carts, err := svc.GetCarts()
		if err != nil {
			svc.Logger.Error(err.Error())
			return
		}
		ch <- carts
	}()

	return ch
}
