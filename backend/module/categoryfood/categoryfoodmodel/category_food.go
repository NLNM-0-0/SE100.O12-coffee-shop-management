package categoryfoodmodel

import (
	"backend/common"
	"backend/module/category/categorymodel"
)

type CategoryFood struct {
	FoodId     string                 `json:"foodId" gorm:"column:foodId;"`
	CategoryId string                 `json:"-" gorm:"column:categoryId;"`
	Category   categorymodel.Category `json:"category" gorm:"foreignkey:CategoryId"`
}

func (*CategoryFood) TableName() string {
	return common.TableCategoryFood
}
