package supplierbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/importnote/importnotemodel"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplier/suppliermodel/filter"
	"context"
)

type SeeSupplierImportNoteRepo interface {
	SeeSupplierImportNote(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) ([]importnotemodel.ImportNote, error)
}

type seeSupplierImportNoteBiz struct {
	repo      SeeSupplierImportNoteRepo
	requester middleware.Requester
}

func NewSeeSupplierImportNoteBiz(
	repo SeeSupplierImportNoteRepo,
	requester middleware.Requester) *seeSupplierImportNoteBiz {
	return &seeSupplierImportNoteBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierImportNoteBiz) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	importNotes, err := biz.repo.SeeSupplierImportNote(
		ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}

	return importNotes, nil
}
