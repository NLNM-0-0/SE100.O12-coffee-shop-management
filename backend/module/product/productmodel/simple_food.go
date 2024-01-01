package productmodel

import "backend/common"

type SimpleFood struct {
	Id   string `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*SimpleFood) TableName() string {
	return common.TableFood
}
