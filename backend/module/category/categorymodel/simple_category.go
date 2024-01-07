package categorymodel

import "backend/common"

type SimpleCategory struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleCategory) TableName() string {
	return common.TableCategory
}
