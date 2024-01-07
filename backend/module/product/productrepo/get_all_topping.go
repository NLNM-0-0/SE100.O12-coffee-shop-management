package productrepo

import (
	"backend/module/product/productmodel"
	"context"
)

type GetAllToppingStore interface {
	GetAllTopping(
		ctx context.Context) ([]productmodel.SimpleTopping, error)
}

type getAllToppingRepo struct {
	store GetAllToppingStore
}

func NewGetAllToppingRepo(store GetAllToppingStore) *getAllToppingRepo {
	return &getAllToppingRepo{store: store}
}

func (repo *getAllToppingRepo) GetAllTopping(
	ctx context.Context) ([]productmodel.SimpleTopping, error) {
	result, err := repo.store.GetAllTopping(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
