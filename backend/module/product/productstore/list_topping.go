package productstore

import (
	"backend/common"
	"backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) ListTopping(
	ctx context.Context,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]productmodel.Topping, error) {
	var result []productmodel.Topping
	db := s.db

	db = db.Table(common.TableTopping)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
