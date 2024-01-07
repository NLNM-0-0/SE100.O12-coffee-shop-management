package ingredientmodel

import (
	"backend/common"
	"backend/module/unittype/unittypemodel"
	"errors"
)

type Ingredient struct {
	Id         string                 `json:"id" gorm:"column:id;"`
	Name       string                 `json:"name" gorm:"column:name;"`
	Amount     float32                `json:"amount" gorm:"column:amount;"`
	UnitTypeId string                 `json:"-" gorm:"column:unitTypeId;"`
	UnitType   unittypemodel.UnitType `json:"unitType"`
	Price      int                    `json:"price" gorm:"column:price;"`
}

func (*Ingredient) TableName() string {
	return common.TableIngredient
}

var (
	ErrIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrIngredientIdInvalid",
	)
	ErrIngredientNameEmpty = common.NewCustomError(
		errors.New("name of ingredient is empty"),
		"Tên của nguyên vật liệu đang trống",
		"ErrIngredientNameEmpty",
	)
	ErrIngredientPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"Giá của nguyên vật liệu đang là số âm",
		"ErrIngredientPriceIsNegativeNumber",
	)
	ErrIngredientUnitTypeInvalid = common.NewCustomError(
		errors.New("unit type of ingredient is invalid"),
		"Đơn vị đo lường của nguyên vật liệu không hợp lệ",
		"ErrIngredientUnitTypeInvalid",
	)
	ErrIngredientIdDuplicate = common.ErrDuplicateKey(
		errors.New("Nguyên vật liệu đã tồn tại"),
	)
	ErrIngredientNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của nguyên vật liệu đã tồn tại"),
	)
	ErrIngredientUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa nguyên vật liệu"),
	)
	ErrIngredientCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền thêm nguyên vật liệu mới"),
	)
	ErrIngredientViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem nguyên vật liệu"),
	)
)
