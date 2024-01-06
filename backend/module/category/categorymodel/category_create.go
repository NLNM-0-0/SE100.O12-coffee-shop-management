package categorymodel

import (
	"backend/common"
)

type CategoryCreate struct {
	Id          *string `json:"id" gorm:"column:id;"`
	Name        string  `json:"name" gorm:"column:name;"`
	Description string  `json:"description" gorm:"column:description;"`
}

func (*CategoryCreate) TableName() string {
	return common.TableCategory
}

func (data *CategoryCreate) Validate() error {
	if !common.ValidateId(data.Id) {
		return ErrCategoryIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
