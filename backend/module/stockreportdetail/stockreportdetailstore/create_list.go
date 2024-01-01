package stockreportdetailstore

import (
	"backend/common"
	"backend/module/stockreportdetail/stockreportdetailmodel"
	"context"
)

func (s *sqlStore) CreateListStockReportDetail(
	ctx context.Context,
	data []stockreportdetailmodel.StockReportDetail) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
