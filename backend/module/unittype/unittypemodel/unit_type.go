package unittypemodel

import (
	"backend/common"
	"errors"
)

type UnitType struct {
	Id          string      `json:"id" gorm:"column:id;"`
	Name        string      `json:"name" gorm:"column:name;"`
	MeasureType MeasureType `json:"measureType" gorm:"column:measureType;"`
	Value       int         `json:"value" gorm:"column:value;"`
}

var (
	ErrUnitTypeMeasureTypeInvalid = common.NewCustomError(
		errors.New("measure type is valid"),
		"Loại đo lường không hợp lệ",
		"ErrUnitTypeMeasureTypeInvalid",
	)
)
