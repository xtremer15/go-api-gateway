package services

import (
	"api-gateway/internal/ports"
	"api-gateway/internal/types/types"
	"api-gateway/pkg/logger"
)

type RecipeService struct {
	Repo   ports.RecipeRepository
	Logger *logger.Logger
}

func NewRecipeService(repo ports.RecipeRepository, logger *logger.Logger) *RecipeService {
	return &RecipeService{
		Repo:   repo,
		Logger: logger,
	}
}

func (svc *RecipeService) GetRecipes() ([]types.Recipe, error) {
	recipes, err := svc.Repo.GetRecipes()
	if err != nil {
		svc.Logger.Error("Failed to fetch recipes")
		return nil, err
	}
	return recipes, nil
}

func (svc *RecipeService) FetchAsyncRecipes() <-chan any {
	ch := make(chan any)

	go func() {
		defer close(ch)
		recipes, err := svc.GetRecipes()
		if err != nil {
			svc.Logger.Error(err.Error())
			return
		}
		ch <- recipes
	}()

	return ch
}
