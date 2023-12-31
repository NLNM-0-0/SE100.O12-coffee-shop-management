package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateStatusTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateStatus) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
