package unittypebiz

import (
	"backend/module/unittype/unittypemodel"
	"context"
)

type ListUnitTypeStore interface {
	ListUnitType(
		ctx context.Context,
		condition map[string]interface{}) ([]unittypemodel.UnitType, error)
}

type getAllUnitType struct {
	store ListUnitTypeStore
}

func NewGetAllUnitTypeBiz(
	store ListUnitTypeStore) *getAllUnitType {
	return &getAllUnitType{
		store: store,
	}
}

func (biz *getAllUnitType) GetAllUnitType(
	ctx context.Context) ([]unittypemodel.UnitType, error) {
	result, err := biz.store.ListUnitType(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result, nil
}
