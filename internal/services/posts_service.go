package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type PostsService struct {
	PostsRepo ports.PostsRepo
	Logger    *logger.Logger
}

func NewPostsService(postsRepo ports.PostsRepo, logger *logger.Logger) *PostsService {
	return &PostsService{
		PostsRepo: postsRepo,
		Logger:    logger,
	}
}

func (svc *PostsService) GetPosts() ([]types.Post, error) {
	posts, err := svc.PostsRepo.GetPosts()
	if err != nil {
		svc.Logger.Error("Posts not found")
		return nil, err
	}
	return posts, nil
}

func (svc *PostsService) FetchAsyncPosts() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)
		posts, err := svc.GetPosts()
		if err != nil {
			svc.Logger.Error(err.Error())
			return
		}
		ch <- posts
	}()

	return ch
}
