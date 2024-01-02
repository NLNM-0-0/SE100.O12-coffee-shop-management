package inventorychecknotedetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"errors"
)

type InventoryCheckNoteDetail struct {
	InventoryCheckNoteId string                           `json:"inventoryCheckNoteId" gorm:"column:inventoryCheckNoteId;"`
	IngredientId         string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient           ingredientmodel.SimpleIngredient `json:"ingredient"`
	Initial              float32                          `json:"initial" gorm:"column:initial;"`
	Difference           float32                          `json:"difference" gorm:"column:difference;"`
	Final                float32                          `json:"final" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetail) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

var (
	ErrInventoryCheckDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrInventoryCheckDetailIngredientIdInvalid",
	)
	ErrInventoryCheckDifferenceIsInvalid = common.NewCustomError(
		errors.New("difference is invalid"),
		"Số lượng chỉnh sửa không hợp lệ",
		"ErrInventoryCheckDifferenceIsInvalid",
	)
)
