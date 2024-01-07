package productmodel

import (
	"backend/common"
	"backend/module/sizefood/sizefoodmodel"
)

type SimpleFoodForSale struct {
	Id             string        `json:"id" gorm:"column:id;"`
	Name           string        `json:"name" gorm:"column:name;"`
	Image          string        `json:"image" gorm:"column:image;"`
	FoodCategories Categories    `json:"categories" gorm:"foreignkey:foodId;association_foreignkey:id"`
	FoodSizes      SizesForSales `json:"sizes" gorm:"foreignkey:foodId;association_foreignkey:id"`
}

type SizesForSales []sizefoodmodel.SizeFoodForSale

func (*SizesForSales) TableName() string {
	return common.TableSizeFood
}
