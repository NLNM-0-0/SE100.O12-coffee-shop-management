package productrepo

import (
	"backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodsStore interface {
	UpdateStatusFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdateStatus,
	) error
}

type changeStatusFoodsRepo struct {
	store ChangeStatusFoodsStore
}

func NewChangeStatusFoodsRepo(store ChangeStatusFoodsStore) *changeStatusFoodsRepo {
	return &changeStatusFoodsRepo{store: store}
}

func (biz *changeStatusFoodsRepo) ChangeStatusFoods(
	ctx context.Context,
	data *productmodel.FoodUpdateStatus) error {
	if err := biz.store.UpdateStatusFood(
		ctx, data.ProductId, data); err != nil {
		return err
	}

	return nil
}
