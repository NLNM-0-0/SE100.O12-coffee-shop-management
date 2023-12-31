package productrepo

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

type ListToppingStore interface {
	ListTopping(
		ctx context.Context,
		filter *productmodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]productmodel.Topping, error)
}

type listToppingRepo struct {
	store ListToppingStore
}

func NewListToppingRepo(store ListToppingStore) *listToppingRepo {
	return &listToppingRepo{store: store}
}

func (repo *listToppingRepo) ListTopping(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging) ([]productmodel.Topping, error) {
	result, err := repo.store.ListTopping(
		ctx,
		filter,
		[]string{"id", "name"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
