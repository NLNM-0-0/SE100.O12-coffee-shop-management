package categorymodel

import "backend/common"

type CategoryUpdateAmountProduct struct {
	AmountProduct *int `json:"amountProduct" gorm:"column:amountProduct;"`
}

func (*CategoryUpdateAmountProduct) TableName() string {
	return common.TableCategory
}
