package productmodel

import "backend/common"

type SimpleTopping struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Price int    `json:"price" gorm:"column:price;"`
}

func (*SimpleTopping) TableName() string {
	return common.TableTopping
}
