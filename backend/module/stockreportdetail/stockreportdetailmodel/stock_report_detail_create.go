package stockreportdetailmodel

import (
	"backend/common"
)

type StockReportDetailCreate struct {
	ReportId     string  `json:"-" gorm:"column:reportId;"`
	IngredientId string  `json:"-" gorm:"column:ingredientId;"`
	Initial      float32 `json:"-" gorm:"column:initial;"`
	Sell         float32 `json:"-" gorm:"column:sell"`
	Import       float32 `json:"-" gorm:"column:import;"`
	Export       float32 `json:"-" gorm:"column:export;"`
	Modify       float32 `json:"-" gorm:"column:modify;"`
	Final        float32 `json:"-" gorm:"column:final;"`
}

func (*StockReportDetailCreate) TableName() string {
	return common.TableStockReportDetail
}
