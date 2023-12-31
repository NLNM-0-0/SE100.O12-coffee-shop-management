package importnotebiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/importnote/importnotemodel"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListImportNoteBySupplierRepo interface {
	ListImportNoteBySupplier(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) ([]importnotemodel.ImportNote, error)
}

type listImportNoteBySupplierBiz struct {
	repo      ListImportNoteBySupplierRepo
	requester middleware.Requester
}

func NewListImportNoteBySupplierBiz(
	repo ListImportNoteBySupplierRepo,
	requester middleware.Requester) *listImportNoteBySupplierBiz {
	return &listImportNoteBySupplierBiz{repo: repo, requester: requester}
}

func (biz *listImportNoteBySupplierBiz) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	result, err := biz.repo.ListImportNoteBySupplier(ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
