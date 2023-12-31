package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateInfo,
) error {

	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return productmodel.ErrToppingNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
