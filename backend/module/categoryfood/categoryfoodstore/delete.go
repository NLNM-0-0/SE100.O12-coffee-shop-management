package categoryfoodstore

import (
	"backend/common"
	"backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) DeleteCategoryFood(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Table(common.TableCategoryFood).
		Where(conditions).
		Delete(&categorymodel.Category{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
