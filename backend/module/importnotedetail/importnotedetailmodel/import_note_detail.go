package importnotedetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
	"errors"
)

type ImportNoteDetail struct {
	ImportNoteId                  string                           `json:"importNoteId" gorm:"column:importNoteId;"`
	IngredientId                  string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient                    ingredientmodel.SimpleIngredient `json:"ingredient" gorm:"foreignKey:IngredientId;references:Id"`
	TempIngredient                *ingredientmodel.Ingredient      `json:"-" gorm:"-"`
	Price                         int                              `json:"price" gorm:"column:price"`
	TotalUnit                     int                              `json:"totalUnit" gorm:"column:totalUnit"`
	AmountImport                  float32                          `json:"amountImport" gorm:"column:amountImport"`
	AmountImportByDefaultUnitType float32                          `json:"-" gorm:"-"`
	UnitTypeName                  string                           `json:"unitTypeName" gorm:"column:unitTypeName;"`
}

func (*ImportNoteDetail) TableName() string {
	return common.TableImportNoteDetail
}

var (
	ErrImportDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrImportDetailIngredientIdInvalid",
	)
	ErrImportDetailPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"Giá của nguyên vật liệu đang là số âm",
		"ErrImportDetailPriceIsNegativeNumber",
	)
	ErrImportDetailAmountImportIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount import is not positive number"),
		"Lượng muốn nhập đang không phải số dương",
		"ErrImportDetailAmountImportIsNotPositiveNumber",
	)
)
