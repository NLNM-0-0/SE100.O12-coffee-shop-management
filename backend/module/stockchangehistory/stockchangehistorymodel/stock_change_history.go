package stockchangehistorymodel

import "backend/common"

type StockChangeHistory struct {
	Id           string           `json:"id" gorm:"column:id;"`
	IngredientId string           `json:"ingredientId" gorm:"column:ingredientId;"`
	Amount       float32          `json:"amount" gorm:"column:amount;"`
	AmountLeft   float32          `json:"amountLeft" gorm:"column:amountLeft;"`
	Type         *StockChangeType `json:"type" gorm:"column:type;"`
}

func (*StockChangeHistory) TableName() string {
	return common.TableStockChangeHistory
}
