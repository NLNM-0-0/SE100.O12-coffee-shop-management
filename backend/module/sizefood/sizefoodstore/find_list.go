package sizefoodstore

import (
	"backend/common"
	"backend/module/sizefood/sizefoodmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindListSizeFood(
	ctx context.Context,
	foodId string) ([]sizefoodmodel.SizeFood, error) {
	var data []sizefoodmodel.SizeFood
	db := s.db

	if err := db.
		Table(common.TableSizeFood).
		Where("foodId = ?", foodId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
