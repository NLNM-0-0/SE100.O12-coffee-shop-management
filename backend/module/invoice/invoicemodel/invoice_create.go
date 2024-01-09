package invoicemodel

import (
	"backend/common"
	"backend/module/invoicedetail/invoicedetailmodel"
)

type InvoiceCreate struct {
	Id                  string                                   `json:"-" gorm:"column:id;"`
	CustomerId          string                                   `json:"customerId" gorm:"column:customerId;"`
	TotalPrice          int                                      `json:"-" gorm:"column:totalPrice;"`
	IsUsePoint          bool                                     `json:"isUsePoint" gorm:"-"`
	TotalCost           int                                      `json:"-" gorm:"column:totalCost;"`
	AmountReceived      int                                      `json:"-" gorm:"column:amountReceived"`
	AmountPriceUsePoint int                                      `json:"-" gorm:"column:amountPriceUsePoint"`
	PointUse            int                                      `json:"-" gorm:"column:pointUse;"`
	PointReceive        int                                      `json:"-" gorm:"column:pointReceive;"`
	CreatedBy           string                                   `json:"-" gorm:"column:createdBy;"`
	InvoiceDetails      []invoicedetailmodel.InvoiceDetailCreate `json:"details" gorm:"-"`
	MapIngredient       map[string]float32                       `json:"-" gorm:"-"`
}

func (*InvoiceCreate) TableName() string {
	return common.TableInvoice
}

func (data *InvoiceCreate) Validate() *common.AppError {
	if !common.ValidateId(&data.CustomerId) {
		return ErrInvoiceCustomerIdInvalid
	}
	if data.InvoiceDetails == nil || len(data.InvoiceDetails) == 0 {
		return ErrInvoiceDetailsEmpty
	}

	for _, invoiceDetail := range data.InvoiceDetails {
		if err := invoiceDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
