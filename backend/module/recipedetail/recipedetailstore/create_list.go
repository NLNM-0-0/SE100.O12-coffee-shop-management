package recipedetailstore

import (
	"backend/common"
	"backend/module/recipedetail/recipedetailmodel"
	"context"
)

func (s *sqlStore) CreateListRecipeDetail(
	ctx context.Context,
	data []recipedetailmodel.RecipeDetailCreate) error {
	db := s.db
	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
