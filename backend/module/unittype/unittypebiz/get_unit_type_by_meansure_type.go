package unittypebiz

import (
	"backend/module/unittype/unittypemodel"
	"context"
)

type getUnitTypeByMeasureType struct {
	store ListUnitTypeStore
}

func NewGetUnitTypeByMeasureTypeBiz(
	store ListUnitTypeStore) *getUnitTypeByMeasureType {
	return &getUnitTypeByMeasureType{
		store: store,
	}
}

func (biz *getUnitTypeByMeasureType) GetUnitTypeByMeasureType(
	ctx context.Context,
	measureType unittypemodel.MeasureType) ([]unittypemodel.UnitType, error) {
	result, err := biz.store.ListUnitType(ctx, map[string]interface{}{"measureType": measureType})
	if err != nil {
		return nil, err
	}
	return result, nil
}
