package ingredientstore

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.Ingredient, error) {
	var result []ingredientmodel.Ingredient
	db := s.db

	db = db.Table(common.TableIngredient)

	if err := db.
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
