package recipedetailmodel

import "backend/common"

type RecipeDetailCreate struct {
	RecipeId     string  `json:"-" gorm:"column:recipeId;"`
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	AmountNeed   float32 `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetailCreate) TableName() string {
	return common.TableRecipeDetail
}

func (data *RecipeDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrRecipeDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumberFloat(data.AmountNeed) {
		return ErrRecipeDetailAmountNeedIsNotPositiveNumber
	}
	return nil
}
