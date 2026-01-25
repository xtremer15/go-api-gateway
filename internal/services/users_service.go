package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type UserService struct {
	UserRepo ports.UserRepo
	Logger   *logger.Logger
}

func NewUserService(userRepo ports.UserRepo, logger *logger.Logger) *UserService {
	return &UserService{
		UserRepo: userRepo,
		Logger:   logger,
	}
}

func (svc *UserService) GetUsers() ([]types.User, error) {
	users, err := svc.UserRepo.GetUsers()

	if err != nil {
		svc.Logger.Error("Users not found")
		return nil, err
	}

	return users, nil
}
