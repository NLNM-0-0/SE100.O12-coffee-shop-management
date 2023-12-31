package sizefoodmodel

import (
	"backend/common"
	"backend/module/recipe/recipemodel"
)

type SizeFoodCreate struct {
	FoodId   string                    `json:"-" gorm:"column:foodId;"`
	SizeId   string                    `json:"-" gorm:"column:sizeId;"`
	Name     string                    `json:"name" gorm:"column:name"`
	Cost     int                       `json:"cost" gorm:"column:cost"`
	Price    int                       `json:"price" gorm:"column:price"`
	RecipeId string                    `json:"-" gorm:"column:recipeId;"`
	Recipe   *recipemodel.RecipeCreate `json:"recipe" gorm:"-"`
}

func (*SizeFoodCreate) TableName() string {
	return common.TableSizeFood
}

func (data *SizeFoodCreate) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrSizeFoodNameEmpty
	}
	if common.ValidateNegativeNumberInt(data.Cost) {
		return ErrSizeFoodCostIsNegativeNumber
	}
	if common.ValidateNegativeNumberInt(data.Price) {
		return ErrSizeFoodPriceIsNegativeNumber
	}
	if data.Recipe == nil {
		return ErrSizeFoodRecipeEmpty
	}
	if err := data.Recipe.Validate(); err != nil {
		return err
	}
	return nil
}
