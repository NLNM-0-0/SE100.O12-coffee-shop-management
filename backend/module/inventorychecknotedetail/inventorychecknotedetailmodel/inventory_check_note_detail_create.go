package inventorychecknotedetailmodel

import "backend/common"

type InventoryCheckNoteDetailCreate struct {
	InventoryCheckNoteId string  `json:"-" gorm:"column:inventoryCheckNoteId;"`
	IngredientId         string  `json:"ingredientId" gorm:"column:ingredientId;"`
	Initial              float32 `json:"-" gorm:"column:initial;"`
	Difference           float32 `json:"difference" gorm:"column:difference;"`
	Final                float32 `json:"-" gorm:"column:final;"`
}

func (*InventoryCheckNoteDetailCreate) TableName() string {
	return common.TableInventoryCheckNoteDetail
}

func (data *InventoryCheckNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrInventoryCheckDetailIngredientIdInvalid
	}
	if data.Difference == 0 {
		return ErrInventoryCheckDifferenceIsInvalid
	}
	return nil
}
