package categoryfoodmodel

import (
	"backend/common"
	"backend/module/category/categorymodel"
)

type CategoryFood struct {
	FoodId     string                       `json:"-" gorm:"column:foodId;"`
	CategoryId string                       `json:"-" gorm:"column:categoryId;"`
	Category   categorymodel.SimpleCategory `json:"category"`
}

func (*CategoryFood) TableName() string {
	return common.TableCategoryFood
}
