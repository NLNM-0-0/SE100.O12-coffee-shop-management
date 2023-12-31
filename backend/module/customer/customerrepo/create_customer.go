package customerrepo

import (
	"backend/module/customer/customermodel"
	"context"
)

type CreateCustomerStore interface {
	CreateCustomer(
		ctx context.Context,
		data *customermodel.CustomerCreate) error
}

type createCustomerRepo struct {
	store CreateCustomerStore
}

func NewCreateCustomerRepo(store CreateCustomerStore) *createCustomerRepo {
	return &createCustomerRepo{store: store}
}

func (repo *createCustomerRepo) CreateCustomer(
	ctx context.Context,
	data *customermodel.CustomerCreate) error {
	if err := repo.store.CreateCustomer(ctx, data); err != nil {
		return err
	}
	return nil
}
