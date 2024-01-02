package ingredientstore

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

func (s *sqlStore) GetAllSimpleIngredient(
	ctx context.Context,
	moreKeys ...string) ([]ingredientmodel.SimpleIngredient, error) {
	var result []ingredientmodel.SimpleIngredient
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
