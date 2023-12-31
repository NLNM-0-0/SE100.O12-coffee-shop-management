package supplierdebtstore

import (
	"backend/common"
	"backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

func (s *sqlStore) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
