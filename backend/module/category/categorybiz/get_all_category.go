package categorybiz

import (
	"backend/module/category/categorymodel"
	"context"
)

type GetAllCategoryStore interface {
	GetAllCategory(
		ctx context.Context) ([]categorymodel.SimpleCategory, error)
}

type getAllCategoryBiz struct {
	store GetAllCategoryStore
}

func NewGetAllCategoryBiz(
	store GetAllCategoryStore) *getAllCategoryBiz {
	return &getAllCategoryBiz{
		store: store,
	}
}

func (biz *getAllCategoryBiz) GetAllCategory(
	ctx context.Context) ([]categorymodel.SimpleCategory, error) {
	result, err := biz.store.GetAllCategory(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
