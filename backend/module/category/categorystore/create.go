package categorystore

import (
	"backend/common"
	"backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) CreateCategory(
	ctx context.Context,
	data *categorymodel.CategoryCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "name"); key {
			case `PRIMARY`:
				return categorymodel.ErrCategoryIdDuplicate
			case "name":
				return categorymodel.ErrCategoryNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
