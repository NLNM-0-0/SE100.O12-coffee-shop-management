package productmodel

import "backend/common"

type ProductsUpdateStatus struct {
	ProductIds []string `json:"Ids" gorm:"-"`
	IsActive   *bool    `json:"isActive" gorm:"column:isActive;"`
}

func (data *ProductsUpdateStatus) Validate() error {
	if data.IsActive == nil {
		return ErrProductStatusEmpty
	}
	for _, v := range data.ProductIds {
		if !common.ValidateNotNilId(&v) {
			return ErrProductIdInvalid
		}
	}
	return nil
}
