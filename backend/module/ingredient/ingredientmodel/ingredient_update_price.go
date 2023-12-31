package ingredientmodel

import "backend/common"

type IngredientUpdatePrice struct {
	Price *float32 `json:"price" gorm:"column:price;"`
}

func (*IngredientUpdatePrice) TableName() string {
	return common.TableIngredient
}
