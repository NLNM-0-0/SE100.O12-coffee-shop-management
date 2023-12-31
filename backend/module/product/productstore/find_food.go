package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindFood(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*productmodel.Food, error) {
	var data productmodel.Food
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
