package importnotedetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
)

type ImportNoteDetailCreate struct {
	ImportNoteId           string                      `json:"-" gorm:"column:importNoteId;"`
	IngredientId           string                      `json:"ingredientId" gorm:"column:ingredientId;"`
	Ingredient             *ingredientmodel.Ingredient `json:"-" gorm:"-"`
	Price                  int                         `json:"price" gorm:"column:price"`
	PriceByDefaultUnitType int                         `json:"-" gorm:"-"`
	IsReplacePrice         bool                        `json:"isReplacePrice" gorm:"-"`
	AmountImport           float32                     `json:"amountImport" json:"amountImport" gorm:"column:amountImport"`
	TotalUnit              int                         `json:"-" gorm:"column:totalUnit"`
	UnitTypeName           string                      `json:"-" gorm:"column:unitTypeName;"`
	UnitTypeId             string                      `json:"unitTypeId" gorm:"-"`
}

func (*ImportNoteDetailCreate) TableName() string {
	return common.TableImportNoteDetail
}

func (data *ImportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrImportDetailIngredientIdInvalid
	}
	if common.ValidateNegativeNumberInt(data.Price) {
		return ErrImportDetailPriceIsNegativeNumber
	}
	if common.ValidateNotPositiveNumberFloat(data.AmountImport) {
		return ErrImportDetailAmountImportIsNotPositiveNumber
	}
	return nil
}
