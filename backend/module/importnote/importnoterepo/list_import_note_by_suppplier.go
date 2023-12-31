package importnoterepo

import (
	"backend/common"
	"backend/module/importnote/importnotemodel"
	"backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListImportNoteBySupplierStore interface {
	ListImportNoteBySupplier(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging,
		moreKeys ...string) ([]importnotemodel.ImportNote, error)
}

type listImportNoteBySupplierRepo struct {
	importNoteStore ListImportNoteBySupplierStore
}

func NewListImportNoteBySupplierRepo(
	importNoteStore ListImportNoteBySupplierStore) *listImportNoteBySupplierRepo {
	return &listImportNoteBySupplierRepo{
		importNoteStore: importNoteStore,
	}
}

func (biz *listImportNoteBySupplierRepo) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	importNotes, err := biz.importNoteStore.ListImportNoteBySupplier(
		ctx, supplierId, filter, paging, "Supplier")
	if err != nil {
		return nil, err
	}

	return importNotes, nil
}
