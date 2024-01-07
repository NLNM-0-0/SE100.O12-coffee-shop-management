package usermodel

import "backend/common"

type UsersUpdateStatus struct {
	UserIds  []string `json:"userIds" gorm:"-"`
	IsActive *bool    `json:"isActive" gorm:"column:isActive;" example:"true"`
}

func (*UsersUpdateStatus) TableName() string {
	return common.TableUser
}

func (data *UsersUpdateStatus) Validate() error {
	if data.IsActive == nil {
		return ErrUserStatusEmpty
	}
	for _, v := range data.UserIds {
		if !common.ValidateNotNilId(&v) {
			return ErrUserIdInvalid
		}
	}
	return nil
}
