package suppliermodel

import "backend/common"

type SimpleSupplier struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Phone string `json:"phone" gorm:"column:phone;"`
}

func (*SimpleSupplier) TableName() string {
	return common.TableSupplier
}
