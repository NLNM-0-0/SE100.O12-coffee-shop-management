package productbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodsRepo interface {
	ChangeStatusFoods(
		ctx context.Context,
		data *productmodel.FoodUpdateStatus) error
}

type changeStatusFoodsBiz struct {
	repo      ChangeStatusFoodsRepo
	requester middleware.Requester
}

func NewChangeStatusFoodsBiz(
	repo ChangeStatusFoodsRepo,
	requester middleware.Requester) *changeStatusFoodsBiz {
	return &changeStatusFoodsBiz{repo: repo, requester: requester}
}

func (biz *changeStatusFoodsBiz) ChangeStatusFoods(
	ctx context.Context,
	data productmodel.ProductsUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.FoodUpdateStatusFeatureCode) {
		return productmodel.ErrFoodChangeStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	for _, v := range data.ProductIds {
		foodUpdateStatus := productmodel.FoodUpdateStatus{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: v,
				IsActive:  data.IsActive,
			},
		}
		if err := biz.repo.ChangeStatusFoods(ctx, &foodUpdateStatus); err != nil {
			return err
		}
	}

	return nil
}
