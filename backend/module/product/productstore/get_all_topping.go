package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) GetAllTopping(
	ctx context.Context) ([]productmodel.SimpleTopping, error) {
	var result []productmodel.SimpleTopping
	db := s.db

	db = db.Table(common.TableTopping)
	db = db.Where("isActive = ?", true)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
