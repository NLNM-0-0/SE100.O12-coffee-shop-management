package stockreportdetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
)

type StockReportDetail struct {
	ReportId     string                           `json:"-" gorm:"column:reportId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient"`
	Initial      int                              `json:"initial" gorm:"column:initial;"`
	Sell         int                              `json:"sell" gorm:"column:sell"`
	Import       int                              `json:"import" gorm:"column:import;"`
	Export       int                              `json:"export" gorm:"column:export;"`
	Modify       int                              `json:"modify" gorm:"column:modify;"`
	Final        int                              `json:"final" gorm:"column:final;"`
}

func (*StockReportDetail) TableName() string {
	return common.TableStockReportDetail
}
