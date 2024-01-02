package stockreportdetailmodel

import (
	"backend/common"
	"backend/module/ingredient/ingredientmodel"
)

type StockReportDetail struct {
	ReportId     string                           `json:"-" gorm:"column:reportId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient"`
	Initial      float32                          `json:"initial" gorm:"column:initial;"`
	Sell         float32                          `json:"sell" gorm:"column:sell"`
	Import       float32                          `json:"import" gorm:"column:import;"`
	Export       float32                          `json:"export" gorm:"column:export;"`
	Modify       float32                          `json:"modify" gorm:"column:modify;"`
	Final        float32                          `json:"final" gorm:"column:final;"`
}

func (*StockReportDetail) TableName() string {
	return common.TableStockReportDetail
}
