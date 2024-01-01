package productmodel

import "backend/common"

type ProductCreate struct {
	Id           *string `json:"id" gorm:"column:id;"`
	Name         string  `json:"name" gorm:"column:name;"`
	Description  string  `json:"description" gorm:"column:description;"`
	CookingGuide string  `json:"cookingGuide" gorm:"column:cookingGuide"`
	Image        string  `json:"image" gorm:"column:image;"`
}

func (data *ProductCreate) Validate() error {
	if !common.ValidateId(data.Id) {
		return ErrProductIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrProductNameEmpty
	}
	if !common.ValidateUrl(data.Image) {
		return ErrProductImageInvalid
	}
	return nil
}
