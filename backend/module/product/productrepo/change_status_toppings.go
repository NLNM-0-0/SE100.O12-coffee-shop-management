package productrepo

import (
	"backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingsStore interface {
	UpdateStatusTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdateStatus,
	) error
}

type changeStatusToppingsRepo struct {
	store ChangeStatusToppingsStore
}

func NewChangeStatusToppingsRepo(store ChangeStatusToppingsStore) *changeStatusToppingsRepo {
	return &changeStatusToppingsRepo{store: store}
}

func (biz *changeStatusToppingsRepo) ChangeStatusToppings(
	ctx context.Context,
	data *productmodel.ToppingUpdateStatus) error {
	if err := biz.store.UpdateStatusTopping(
		ctx, data.ProductId, data); err != nil {
		return err
	}

	return nil
}
