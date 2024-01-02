package ingredientmodel

import "backend/common"

type IngredientUpdateAmount struct {
	Amount float32 `json:"amount" gorm:"-"`
}

func (*IngredientUpdateAmount) TableName() string {
	return common.TableIngredient
}
