package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) CreateFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "name"); key {
			case "PRIMARY":
				return productmodel.ErrFoodIdDuplicate
			case "name":
				return productmodel.ErrFoodNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
