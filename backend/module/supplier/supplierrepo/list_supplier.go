package supplierrepo

import (
	"backend/common"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListSupplierStore interface {
	ListSupplier(
		ctx context.Context,
		filter *filter.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]suppliermodel.Supplier, error)
}

type listSupplierRepo struct {
	store ListSupplierStore
}

func NewListSupplierRepo(store ListSupplierStore) *listSupplierRepo {
	return &listSupplierRepo{store: store}
}

func (repo *listSupplierRepo) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	result, err := repo.store.ListSupplier(
		ctx,
		filter,
		[]string{"id", "name", "email", "phone"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
