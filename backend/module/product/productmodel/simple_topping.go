package productmodel

import "backend/common"

type SimpleTopping struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleTopping) TableName() string {
	return common.TableTopping
}
