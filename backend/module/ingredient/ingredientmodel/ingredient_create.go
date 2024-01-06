package ingredientmodel

import (
	"backend/common"
)

type IngredientCreate struct {
	Id         *string `json:"id" gorm:"column:id;"`
	Name       string  `json:"name" gorm:"column:name;"`
	UnitTypeId string  `json:"unitTypeId" gorm:"column:unitTypeId;"`
	Price      int     `json:"price" gorm:"column:price;"`
}

func (*IngredientCreate) TableName() string {
	return common.TableIngredient
}

func (data *IngredientCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrIngredientIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrIngredientNameEmpty
	}
	if !common.ValidateNotNilId(&data.UnitTypeId) {
		return ErrIngredientUnitTypeInvalid
	}
	if common.ValidateNegativeNumberInt(data.Price) {
		return ErrIngredientPriceIsNegativeNumber
	}
	return nil
}
