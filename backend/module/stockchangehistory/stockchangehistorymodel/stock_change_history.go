package stockchangehistorymodel

import "backend/common"

type StockChangeHistory struct {
	Id           string           `json:"id" gorm:"column:id;"`
	IngredientId string           `json:"ingredientId" gorm:"column:ingredientId;"`
	Amount       int              `json:"amount" gorm:"column:amount;"`
	AmountLeft   int              `json:"amountLeft" gorm:"column:amountLeft;"`
	Type         *StockChangeType `json:"type" gorm:"column:type;"`
}

func (*StockChangeHistory) TableName() string {
	return common.TableStockChangeHistory
}
