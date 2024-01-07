package productrepo

import (
	"backend/module/product/productmodel"
	"context"
)

type GetAllFoodStore interface {
	GetAllFood(
		ctx context.Context,
		moreKeys ...string) ([]productmodel.SimpleFoodForSale, error)
}

type getAllFoodRepo struct {
	store GetAllFoodStore
}

func NewGetAllFoodRepo(store GetAllFoodStore) *getAllFoodRepo {
	return &getAllFoodRepo{store: store}
}

func (repo *getAllFoodRepo) GetAllFood(
	ctx context.Context) ([]productmodel.SimpleFoodForSale, error) {
	result, err := repo.store.GetAllFood(
		ctx,
		"FoodCategories.Category", "FoodSizes")

	if err != nil {
		return nil, err
	}

	return result, nil
}
