package productmodel

import "backend/common"

type FoodWithAmount struct {
	Id     string `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Amount int    `json:"amount" gorm:"-"`
}

func (*FoodWithAmount) TableName() string {
	return common.TableFood
}
