package unittypestore

import (
	"backend/common"
	"backend/module/unittype/unittypemodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUnitType(
	ctx context.Context,
	conditions map[string]interface{}) (*unittypemodel.UnitType, error) {
	var data unittypemodel.UnitType
	db := s.db

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
