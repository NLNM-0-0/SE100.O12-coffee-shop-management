package recipedetailstore

import (
	"backend/common"
	"backend/module/recipedetail/recipedetailmodel"
	"context"
)

func (s *sqlStore) DeleteRecipeDetail(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Where(conditions).
		Delete(recipedetailmodel.RecipeDetail{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
