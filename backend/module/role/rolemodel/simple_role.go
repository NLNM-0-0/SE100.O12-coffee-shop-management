package rolemodel

import "backend/common"

type SimpleRole struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleRole) TableName() string {
	return common.TableRole
}
