package ingredientmodel

import (
	"backend/common"
	"backend/module/unittype/unittypemodel"
)

type SimpleIngredient struct {
	Id         string                 `json:"id" gorm:"column:id;"`
	Name       string                 `json:"name" gorm:"column:name;"`
	UnitTypeId string                 `json:"-" gorm:"column:unitTypeId;"`
	UnitType   unittypemodel.UnitType `json:"unitType,omitempty"`
}

func (*SimpleIngredient) TableName() string {
	return common.TableIngredient
}
