package ingredientbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/ingredient/ingredientmodel"
	"context"
)

type ListIngredientStore interface {
	ListIngredient(
		ctx context.Context,
		filter *ingredientmodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ingredientmodel.Ingredient, error)
}

type listIngredientBiz struct {
	store     ListIngredientStore
	requester middleware.Requester
}

func NewListIngredientBiz(
	store ListIngredientStore,
	requester middleware.Requester) *listIngredientBiz {
	return &listIngredientBiz{store: store, requester: requester}
}

func (biz *listIngredientBiz) ListIngredient(
	ctx context.Context,
	filter *ingredientmodel.Filter,
	paging *common.Paging) ([]ingredientmodel.Ingredient, error) {
	if !biz.requester.IsHasFeature(common.IngredientViewFeatureCode) {
		return nil, ingredientmodel.ErrIngredientViewNoPermission
	}

	result, err := biz.store.ListIngredient(
		ctx,
		filter,
		[]string{"id", "name"},
		paging,
		"UnitType",
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
