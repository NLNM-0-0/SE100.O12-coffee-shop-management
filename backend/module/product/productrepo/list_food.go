package productrepo

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

type ListFoodStore interface {
	ListFood(
		ctx context.Context,
		filter *productmodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]productmodel.Food, error)
}

type listFoodRepo struct {
	store ListFoodStore
}

func NewListFoodRepo(store ListFoodStore) *listFoodRepo {
	return &listFoodRepo{store: store}
}

func (repo *listFoodRepo) ListFood(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging) ([]productmodel.Food, error) {
	result, err := repo.store.ListFood(
		ctx,
		filter,
		[]string{"id", "name"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
