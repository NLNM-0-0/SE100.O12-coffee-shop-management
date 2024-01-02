package ingredientbiz

import (
	"backend/module/ingredient/ingredientmodel"
	"context"
)

type GetAllIngredientStore interface {
	GetAllIngredient(
		ctx context.Context,
		moreKeys ...string) ([]ingredientmodel.Ingredient, error)
}

type getAllIngredientBiz struct {
	store GetAllIngredientStore
}

func NewGetAllIngredientBiz(
	store GetAllIngredientStore) *getAllIngredientBiz {
	return &getAllIngredientBiz{store: store}
}

func (biz *getAllIngredientBiz) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.Ingredient, error) {
	result, err := biz.store.GetAllIngredient(ctx, "UnitType")
	if err != nil {
		return nil, err
	}
	return result, nil
}
