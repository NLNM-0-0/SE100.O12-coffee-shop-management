package productmodel

import (
	"backend/common"
	"errors"
)

type Product struct {
	Id           string `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Description  string `json:"description" gorm:"column:description;"`
	CookingGuide string `json:"cookingGuide" gorm:"column:cookingGuide;"`
	Image        string `json:"image" gorm:"column:image;"`
	IsActive     bool   `json:"isActive" gorm:"column:isActive;"`
}

var (
	ErrProductIdInvalid = common.NewCustomError(
		errors.New("id of product is invalid"),
		"Mã không hợp lệ",
		"ErrProductIdInvalid",
	)
	ErrProductNameEmpty = common.NewCustomError(
		errors.New("name of product is empty"),
		"Tên đang trống",
		"ErrProductNameEmpty",
	)
	ErrProductImageInvalid = common.NewCustomError(
		errors.New("image of product is empty"),
		"Ảnh không hợp lệ",
		"ErrProductImageInvalid",
	)
	ErrProductIsActiveEmpty = common.NewCustomError(
		errors.New("status of product is empty"),
		"Trạng thái đang trống",
		"ErrProductIsActiveEmpty",
	)
	ErrProductInactive = common.NewCustomError(
		errors.New("product has been inactive"),
		"Đã ngừng bán",
		"ErrProductInactive",
	)
)
