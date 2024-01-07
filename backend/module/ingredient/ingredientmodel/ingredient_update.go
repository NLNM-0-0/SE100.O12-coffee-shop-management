package ingredientmodel

import "backend/common"

type IngredientUpdate struct {
	Name       *string `json:"name" gorm:"column:name;"`
	UnitTypeId *string `json:"unitTypeId" gorm:"column:unitTypeId;"`
	Price      *int    `json:"price" gorm:"column:price;"`
}

func (*IngredientUpdate) TableName() string {
	return common.TableIngredient
}

func (data *IngredientUpdate) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrIngredientNameEmpty
	}
	if data.UnitTypeId != nil && !common.ValidateNotNilId(data.UnitTypeId) {
		return ErrIngredientUnitTypeInvalid
	}
	if data.Price != nil && common.ValidateNegativeNumberInt(*data.Price) {
		return ErrIngredientPriceIsNegativeNumber
	}
	return nil
}
