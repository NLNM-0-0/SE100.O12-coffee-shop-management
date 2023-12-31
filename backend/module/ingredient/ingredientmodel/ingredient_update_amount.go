package ingredientmodel

import "backend/common"

type IngredientUpdateAmount struct {
	Amount int `json:"amount" gorm:"-"`
}

func (*IngredientUpdateAmount) TableName() string {
	return common.TableIngredient
}
