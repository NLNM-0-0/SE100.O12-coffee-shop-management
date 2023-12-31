package invoicestore

import (
	"backend/common"
	"backend/module/invoice/invoicemodel"
	"context"
)

func (s *sqlStore) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
