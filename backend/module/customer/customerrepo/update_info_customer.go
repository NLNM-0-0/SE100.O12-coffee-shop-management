package customerrepo

import (
	"backend/module/customer/customermodel"
	"context"
)

type UpdateInfoCustomerStore interface {
	UpdateCustomerInfo(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdateInfo,
	) error
}

type updateInfoCustomerRepo struct {
	store UpdateInfoCustomerStore
}

func NewUpdateInfoCustomerRepo(store UpdateInfoCustomerStore) *updateInfoCustomerRepo {
	return &updateInfoCustomerRepo{store: store}
}

func (repo *updateInfoCustomerRepo) UpdateCustomerInfo(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateInfo) error {
	if err := repo.store.UpdateCustomerInfo(ctx, customerId, data); err != nil {
		return err
	}
	return nil
}
