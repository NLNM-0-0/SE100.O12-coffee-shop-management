package usermodel

import (
	"backend/common"
)

type UserCreate struct {
	Id       *string `json:"id" gorm:"column:id;"`
	Name     string  `json:"name" gorm:"column:name;"`
	Phone    string  `json:"phone" gorm:"column:phone;"`
	Address  string  `json:"address" gorm:"column:address;"`
	Email    string  `json:"email" gorm:"column:email;"`
	Password string  `json:"-" gorm:"column:password;"`
	Salt     string  `json:"-" gorm:"column:salt;"`
	Image    string  `json:"image" gorm:"column:image;"`
	RoleId   string  `json:"roleId" gorm:"column:roleId;"`
}

func (*UserCreate) TableName() string {
	return common.TableUser
}

func (data *UserCreate) Validate() error {
	if !common.ValidateId(data.Id) {
		return ErrUserIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrUserNameEmpty
	}
	if !common.ValidateEmail(data.Email) {
		return ErrUserEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrUserPhoneInvalid
	}
	if !common.ValidateNotNilId(&data.RoleId) {
		return ErrUserRoleInvalid
	}
	if !common.ValidateUrl(data.Image) {
		return ErrUserImageInvalid
	}
	return nil
}
