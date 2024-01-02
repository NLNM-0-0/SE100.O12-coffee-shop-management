package exportnotedetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
)

type ExportNoteDetailCreate struct {
	ExportNoteId                  string                      `json:"-" gorm:"column:exportNoteId;"`
	IngredientId                  string                      `json:"ingredientId" gorm:"column:ingredientId;"`
	Ingredient                    *ingredientmodel.Ingredient `json:"-" gorm:"-"`
	AmountExport                  float32                     `json:"amountExport" gorm:"column:amountExport;"`
	AmountExportByDefaultUnitType float32                     `json:"-" gorm:"-"`
	UnitTypeName                  string                      `json:"-" gorm:"column:unitTypeName;"`
	UnitTypeId                    string                      `json:"unitTypeId" gorm:"-"`
}

func (*ExportNoteDetailCreate) TableName() string {
	return common.TableExportNoteDetail
}

func (data *ExportNoteDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrExportDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumberFloat(data.AmountExport) {
		return ErrExportDetailAmountExportIsNotPositiveNumber
	}
	return nil
}
