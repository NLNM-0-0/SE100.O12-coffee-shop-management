package supplierbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/suppliermodel/filter"
	"backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type SeeSupplierDebtRepo interface {
	SeeSupplierDebt(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierDebtFilter,
		paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error)
}

type seeSupplierDebtBiz struct {
	repo      SeeSupplierDebtRepo
	requester middleware.Requester
}

func NewSeeSupplierDebtBiz(
	repo SeeSupplierDebtRepo,
	requester middleware.Requester) *seeSupplierDebtBiz {
	return &seeSupplierDebtBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierDebtBiz) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	supplierDebts, err := biz.repo.SeeSupplierDebt(
		ctx, supplierId, filterSupplierDebt, paging)
	if err != nil {
		return nil, err
	}

	return supplierDebts, nil
}
