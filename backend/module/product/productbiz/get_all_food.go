package productbiz

import (
	"backend/module/product/productmodel"
	"context"
)

type GetAllFoodRepo interface {
	GetAllFood(
		ctx context.Context) ([]productmodel.SimpleFoodForSale, error)
}

type getAllFoodBiz struct {
	repo GetAllFoodRepo
}

func NewGetAllFoodBiz(
	repo GetAllFoodRepo) *getAllFoodBiz {
	return &getAllFoodBiz{repo: repo}
}

func (biz *getAllFoodBiz) GetAllFood(
	ctx context.Context) ([]productmodel.SimpleFoodForSale, error) {
	result, err := biz.repo.GetAllFood(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
