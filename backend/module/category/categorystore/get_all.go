package categorystore

import (
	"backend/common"
	"backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) GetAllCategory(
	ctx context.Context) ([]categorymodel.SimpleCategory, error) {
	var result []categorymodel.SimpleCategory
	db := s.db

	db = db.Table(common.TableCategory)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
