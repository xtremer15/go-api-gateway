package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type CommentService struct {
	CommentsRepo ports.CommentRepo
	Logger       *logger.Logger
}

func (svc *CommentService) GetComments() ([]types.Comment, error) {
	comments, err := svc.CommentsRepo.GetComments()

	if err != nil {
		svc.Logger.Error("Comments not found")
		return nil, err
	}

	return comments, nil
}
