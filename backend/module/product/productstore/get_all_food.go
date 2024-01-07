package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) GetAllFood(
	ctx context.Context,
	moreKeys ...string) ([]productmodel.SimpleFoodForSale, error) {
	var result []productmodel.SimpleFoodForSale
	db := s.db

	db = db.Table(common.TableFood)
	db = db.Where("isActive = ?", true)

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
