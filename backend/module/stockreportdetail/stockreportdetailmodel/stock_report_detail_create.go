package stockreportdetailmodel

import (
	"backend/common"
)

type StockReportDetailCreate struct {
	ReportId     string `json:"-" gorm:"column:reportId;"`
	IngredientId string `json:"-" gorm:"column:ingredientId;"`
	Initial      int    `json:"-" gorm:"column:initial;"`
	Sell         int    `json:"-" gorm:"column:sell"`
	Import       int    `json:"-" gorm:"column:import;"`
	Export       int    `json:"-" gorm:"column:export;"`
	Modify       int    `json:"-" gorm:"column:modify;"`
	Final        int    `json:"-" gorm:"column:final;"`
}

func (*StockReportDetailCreate) TableName() string {
	return common.TableStockReportDetail
}
