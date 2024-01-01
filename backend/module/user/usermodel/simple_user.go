package usermodel

import "backend/common"

type SimpleUser struct {
	Id    string `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Image string `json:"image" gorm:"column:image;"`
	Email string `json:"email" gorm:"column:email;"`
}

func (*SimpleUser) TableName() string {
	return common.TableUser
}
