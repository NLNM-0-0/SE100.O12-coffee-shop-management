package sizefoodmodel

import "backend/common"

type SizeFoodForSale struct {
	FoodId string `json:"-" gorm:"column:foodId;"`
	SizeId string `json:"sizeId" gorm:"column:sizeId;"`
	Name   string `json:"name" gorm:"column:name;"`
	Price  int    `json:"price" gorm:"column:price;"`
}

func (*SizeFoodForSale) TableName() string {
	return common.TableSizeFood
}
