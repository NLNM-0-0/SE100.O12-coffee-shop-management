package invoicestore

import (
	"backend/common"
	"backend/module/invoice/invoicemodel"
	"context"
)

type InvoiceCreateTemp struct {
	Id                  string  `json:"-" gorm:"column:id;"`
	CustomerId          *string `json:"customerId" gorm:"column:customerId;"`
	TotalPrice          int     `json:"-" gorm:"column:totalPrice;"`
	TotalCost           int     `json:"-" gorm:"column:totalCost;"`
	AmountReceived      int     `json:"-" gorm:"column:amountReceived"`
	AmountPriceUsePoint int     `json:"-" gorm:"column:amountPriceUsePoint"`
	PointUse            int     `json:"-" gorm:"column:pointUse;"`
	PointReceive        int     `json:"-" gorm:"column:pointReceive;"`
	CreatedBy           string  `json:"-" gorm:"column:createdBy;"`
}

func (*InvoiceCreateTemp) TableName() string {
	return common.TableInvoice
}

func (s *sqlStore) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	db := s.db

	var customerId *string
	if data.CustomerId != "" {
		customerId = &data.CustomerId
	}
	temp := InvoiceCreateTemp{
		Id:                  data.Id,
		CustomerId:          customerId,
		TotalPrice:          data.TotalPrice,
		TotalCost:           data.TotalCost,
		AmountReceived:      data.AmountReceived,
		AmountPriceUsePoint: data.AmountPriceUsePoint,
		PointUse:            data.PointUse,
		PointReceive:        data.PointReceive,
		CreatedBy:           data.CreatedBy,
	}
	if err := db.Create(temp).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
