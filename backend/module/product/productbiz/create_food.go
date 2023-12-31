package productbiz

import (
	"backend/common"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/product/productmodel"
	"context"
)

type CreateFoodRepo interface {
	CreateFood(
		ctx context.Context,
		data *productmodel.FoodCreate,
	) error
	HandleCategoryFood(
		ctx context.Context,
		foodId string,
		data *productmodel.FoodCreate,
	) error
	HandleSizeFood(
		ctx context.Context,
		data *productmodel.FoodCreate,
	) error
}

type createFoodBiz struct {
	gen       generator.IdGenerator
	repo      CreateFoodRepo
	requester middleware.Requester
}

func NewCreateFoodBiz(
	gen generator.IdGenerator,
	repo CreateFoodRepo,
	requester middleware.Requester) *createFoodBiz {
	return &createFoodBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createFoodBiz) CreateFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	if !biz.requester.IsHasFeature(common.FoodCreateFeatureCode) {
		return productmodel.ErrFoodCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleFoodId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateFood(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleCategoryFood(ctx, *data.Id, data); err != nil {
		return err
	}

	if err := biz.repo.HandleSizeFood(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleFoodId(gen generator.IdGenerator, data *productmodel.FoodCreate) error {
	//handle id for food
	id, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = id

	//handle id for size food
	for i, _ := range data.Sizes {
		data.Sizes[i].FoodId = *id

		sizeId, err := gen.GenerateId()
		if err != nil {
			return err
		}
		data.Sizes[i].SizeId = sizeId

		recipeId, err := gen.GenerateId()
		if err != nil {
			return err
		}
		data.Sizes[i].RecipeId = recipeId
		data.Sizes[i].Recipe.Id = recipeId
		for j, _ := range data.Sizes[i].Recipe.Details {
			data.Sizes[i].Recipe.Details[j].RecipeId = recipeId
		}
	}

	return nil
}
