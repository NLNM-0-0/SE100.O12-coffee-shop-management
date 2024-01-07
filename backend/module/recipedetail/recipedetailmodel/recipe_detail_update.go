package recipedetailmodel

import "backend/common"

type RecipeDetailUpdate struct {
	IngredientId string  `json:"ingredientId" gorm:"-"`
	AmountNeed   float32 `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetailUpdate) TableName() string {
	return common.TableRecipeDetail
}

func (data *RecipeDetailUpdate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrRecipeDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumberFloat(data.AmountNeed) {
		return ErrRecipeDetailAmountNeedIsNotPositiveNumber
	}
	return nil
}
