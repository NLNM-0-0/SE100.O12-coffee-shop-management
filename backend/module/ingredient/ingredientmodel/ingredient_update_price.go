package ingredientmodel

import "backend/common"

type IngredientUpdatePrice struct {
	Price *int `json:"price" gorm:"column:price;"`
}

func (*IngredientUpdatePrice) TableName() string {
	return common.TableIngredient
}
