package ingredientstore

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) UpdateIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdate) error {
	db := s.db

	if err := db.Table(common.TableIngredient).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
