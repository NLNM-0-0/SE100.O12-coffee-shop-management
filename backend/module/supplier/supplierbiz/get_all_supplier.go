package supplierbiz

import (
	"backend/module/supplier/suppliermodel"
	"context"
)

type GetAllSupplierRepo interface {
	GetAllSupplier(
		ctx context.Context) ([]suppliermodel.SimpleSupplier, error)
}

type getAllSupplierBiz struct {
	repo GetAllSupplierRepo
}

func NewGetAllSupplierBiz(
	repo GetAllSupplierRepo) *getAllSupplierBiz {
	return &getAllSupplierBiz{repo: repo}
}

func (biz *getAllSupplierBiz) GetAllSupplier(
	ctx context.Context) ([]suppliermodel.SimpleSupplier, error) {
	result, err := biz.repo.GetAllSupplier(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
