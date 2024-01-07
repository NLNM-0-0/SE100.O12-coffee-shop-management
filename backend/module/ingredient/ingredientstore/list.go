package ingredientstore

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListIngredient(
	ctx context.Context,
	filter *ingredientmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]ingredientmodel.Ingredient, error) {
	var result []ingredientmodel.Ingredient
	db := s.db

	db = db.Table(common.TableIngredient)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

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

func handleFilter(
	db *gorm.DB,
	filter *ingredientmodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.UnitType != "" {
			db = db.Where("unitTypeId = ?", filter.UnitType)
		}
		if filter.MinPrice != nil {
			db = db.Where("price >= ?", filter.MinPrice)
		}
		if filter.MaxPrice != nil {
			db = db.Where("price <= ?", filter.MaxPrice)
		}
		if filter.MinAmount != nil {
			db = db.Where("amount >= ?", filter.MinAmount)
		}
		if filter.MaxAmount != nil {
			db = db.Where("amount <= ?", filter.MaxAmount)
		}
	}
}
