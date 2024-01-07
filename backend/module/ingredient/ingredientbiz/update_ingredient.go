package ingredientbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

type UpdateInfoIngredientRepo interface {
	UpdateIngredient(
		ctx context.Context,
		ingredientId string,
		data *ingredientmodel.IngredientUpdate,
	) error
}

type updateInfoIngredientBiz struct {
	repo      UpdateInfoIngredientRepo
	requester middleware.Requester
}

func NewUpdateInfoIngredientBiz(
	repo UpdateInfoIngredientRepo,
	requester middleware.Requester) *updateInfoIngredientBiz {
	return &updateInfoIngredientBiz{repo: repo, requester: requester}
}

func (biz *updateInfoIngredientBiz) UpdateInfoIngredient(
	ctx context.Context,
	id string,
	data *ingredientmodel.IngredientUpdate) error {
	if !biz.requester.IsHasFeature(common.IngredientUpdateFeatureCode) {
		return ingredientmodel.ErrIngredientUpdateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateIngredient(ctx, id, data); err != nil {
		return err
	}

	return nil
}
