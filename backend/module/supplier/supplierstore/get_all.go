package supplierstore

import (
	"backend/common"
	"backend/module/supplier/suppliermodel"
	"context"
)

func (s *sqlStore) GetAllSupplier(
	ctx context.Context,
	moreKeys ...string) ([]suppliermodel.SimpleSupplier, error) {
	var result []suppliermodel.SimpleSupplier
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
