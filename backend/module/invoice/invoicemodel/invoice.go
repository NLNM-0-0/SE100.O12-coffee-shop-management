package invoicemodel

import (
	"backend/common"
	"backend/module/invoicedetail/invoicedetailmodel"
	"backend/module/user/usermodel"
	"errors"
	"time"
)

type Invoice struct {
	Id                  string                             `json:"id" gorm:"column:id;"`
	CustomerId          string                             `json:"-" gorm:"column:customerId;"`
	Customer            *SimpleCustomer                    `json:"customer"  gorm:"foreignKey:CustomerId;references:Id"`
	TotalPrice          int                                `json:"totalPrice" gorm:"column:totalPrice;"`
	AmountReceived      int                                `json:"amountReceived" gorm:"column:amountReceived"`
	AmountPriceUsePoint int                                `json:"amountPriceUsePoint" gorm:"column:amountPriceUsePoint"`
	PointUse            int                                `json:"pointUse" gorm:"column:pointUse;"`
	PointReceive        int                                `json:"pointReceive" gorm:"column:pointReceive;"`
	CreatedBy           string                             `json:"-" gorm:"column:createdBy;"`
	CreatedByUser       usermodel.SimpleUser               `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	CreatedAt           *time.Time                         `json:"createdAt" gorm:"column:createdAt;"`
	Details             []invoicedetailmodel.InvoiceDetail `json:"details,omitempty"`
}

func (*Invoice) TableName() string {
	return common.TableInvoice
}

var (
	ErrInvoiceCustomerIdInvalid = common.NewCustomError(
		errors.New("id of customer is invalid"),
		"Khách hàng không hợp lệ",
		"ErrInvoiceCustomerIdInvalid",
	)
	ErrInvoiceNotHaveCustomerToUsePoint = common.NewCustomError(
		errors.New("customer is required for this invoice"),
		"Khách hàng bắt buộc phải có cho hóa đơn này",
		"ErrInvoiceNotHaveCustomerToUsePoint",
	)
	ErrInvoiceDetailsEmpty = common.NewCustomError(
		errors.New("list invoice details are empty"),
		"Danh sách sản phẩm cần thanh toán đang trống",
		"ErrInvoiceDetailsEmpty",
	)
	ErrInvoiceIngredientIsNotEnough = common.NewCustomError(
		errors.New("exist ingredient in the stock is not enough for the invoice"),
		"Tồn tại 1 nguyên vật liệu có số lượng trong kho không đủ để bán",
		"ErrInvoiceIngredientIsNotEnough",
	)
	ErrInvoiceCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo hóa đơn"),
	)
	ErrInvoiceViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem hóa đơn"),
	)
)
