package productmodel

import "backend/common"

type ProductUpdateInfo struct {
	Name         *string `json:"name" gorm:"column:name;"`
	Description  *string `json:"description" gorm:"column:description;"`
	CookingGuide *string `json:"cookingGuide" gorm:"column:cookingGuide"`
	Image        *string `json:"image" gorm:"column:image;"`
}

func (data *ProductUpdateInfo) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrProductNameEmpty
	}
	if data.Image != nil && !common.ValidateUrl(*data.Image) {
		return ErrProductImageInvalid
	}
	return nil
}
