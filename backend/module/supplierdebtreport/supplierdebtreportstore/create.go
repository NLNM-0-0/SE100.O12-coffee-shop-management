package supplierdebtreportstore

import (
	"backend/common"
	"backend/module/supplierdebtreport/supplierdebtreportmodel"
	"context"
)

func (s *sqlStore) CreateSupplierDebtReport(
	ctx context.Context,
	data *supplierdebtreportmodel.ReqFindSupplierDebtReport) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
