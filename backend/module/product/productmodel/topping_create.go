package productmodel

import (
	"backend/common"
	"backend/module/recipe/recipemodel"
)

type ToppingCreate struct {
	*ProductCreate `json:",inline"`
	Cost           int                       `json:"cost" gorm:"column:cost;"`
	Price          int                       `json:"price" gorm:"column:price;"`
	RecipeId       string                    `json:"-" gorm:"column:recipeId;"`
	Recipe         *recipemodel.RecipeCreate `json:"recipe" gorm:"-"`
}

func (*ToppingCreate) TableName() string {
	return common.TableTopping
}

func (data *ToppingCreate) Validate() error {
	if data.ProductCreate == nil {
		return ErrToppingProductInfoEmpty
	}
	if err := (*data.ProductCreate).Validate(); err != nil {
		return err
	}
	if common.ValidateNegativeNumberInt(data.Cost) {
		return ErrToppingCostIsNegativeNumber
	}
	if common.ValidateNegativeNumberInt(data.Price) {
		return ErrToppingPriceIsNegativeNumber
	}
	if data.Recipe == nil {
		return ErrToppingRecipeEmpty
	}
	if err := data.Recipe.Validate(); err != nil {
		return err
	}
	return nil
}
