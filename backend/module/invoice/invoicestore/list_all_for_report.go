package invoicestore

import (
	"backend/common"
	"backend/module/invoice/invoicemodel"
	"context"
	"time"
)

func (s *sqlStore) ListAllInvoiceForReport(
	ctx context.Context,
	startTime time.Time,
	endTime time.Time,
	moreKeys ...string) ([]invoicemodel.Invoice, error) {
	var result []invoicemodel.Invoice
	db := s.db

	db = db.Table(common.TableInvoice)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Where("createdAt >= ?", startTime).
		Where("createdAt <= ?", endTime).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
