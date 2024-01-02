package ingredientstore

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) GetAllIngredient(
	ctx context.Context,
	moreKeys ...string) ([]ingredientmodel.Ingredient, error) {
	var result []ingredientmodel.Ingredient
	db := s.db

	db = db.Table(common.TableIngredient)

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
