package productbiz

import (
	"backend/middleware"
	"backend/module/product/productmodel"
	"context"
)

type GetAllToppingRepo interface {
	GetAllTopping(
		ctx context.Context) ([]productmodel.SimpleTopping, error)
}

type getAllToppingBiz struct {
	repo      GetAllToppingRepo
	requester middleware.Requester
}

func NewGetAllToppingBiz(
	repo GetAllToppingRepo) *getAllToppingBiz {
	return &getAllToppingBiz{repo: repo}
}

func (biz *getAllToppingBiz) GetAllTopping(
	ctx context.Context) ([]productmodel.SimpleTopping, error) {
	result, err := biz.repo.GetAllTopping(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
