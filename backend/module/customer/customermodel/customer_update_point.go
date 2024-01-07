package customermodel

import "backend/common"

type CustomerUpdatePoint struct {
	Amount *int `json:"amount" gorm:"-"`
}

func (*CustomerUpdatePoint) TableName() string {
	return common.TableCustomer
}
