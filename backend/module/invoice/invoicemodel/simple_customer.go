package invoicemodel

import "backend/common"

type SimpleCustomer struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Phone string `json:"phone" gorm:"column:phone;"`
	Point int    `json:"point" gorm:"column:point;"`
}

func (*SimpleCustomer) TableName() string {
	return common.TableCustomer
}
