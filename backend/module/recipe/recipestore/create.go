package recipestore

import (
	"backend/common"
	"backend/module/recipe/recipemodel"
	"context"
)

func (s *sqlStore) CreateRecipe(
	ctx context.Context,
	data *recipemodel.RecipeCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
