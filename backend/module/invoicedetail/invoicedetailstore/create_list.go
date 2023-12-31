package invoicedetailstore

import (
	"backend/common"
	"backend/module/invoicedetail/invoicedetailmodel"
	"context"
)

func (s *sqlStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []invoicedetailmodel.InvoiceDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
