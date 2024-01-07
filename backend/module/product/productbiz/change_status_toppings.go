package productbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingsRepo interface {
	ChangeStatusToppings(
		ctx context.Context,
		data *productmodel.ToppingUpdateStatus) error
}

type changeStatusToppingsBiz struct {
	repo      ChangeStatusToppingsRepo
	requester middleware.Requester
}

func NewChangeStatusToppingsBiz(
	repo ChangeStatusToppingsRepo,
	requester middleware.Requester) *changeStatusToppingsBiz {
	return &changeStatusToppingsBiz{repo: repo, requester: requester}
}

func (biz *changeStatusToppingsBiz) ChangeStatusToppings(
	ctx context.Context,
	data productmodel.ProductsUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.ToppingUpdateStatusFeatureCode) {
		return productmodel.ErrToppingChangeStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	for _, v := range data.ProductIds {
		toppingUpdateStatus := productmodel.ToppingUpdateStatus{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: v,
				IsActive:  data.IsActive,
			},
		}
		if err := biz.repo.ChangeStatusToppings(ctx, &toppingUpdateStatus); err != nil {
			return err
		}
	}

	return nil
}
