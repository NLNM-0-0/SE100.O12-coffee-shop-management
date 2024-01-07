package invoicebiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/invoice/invoicemodel"
	"backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

type SeeShopGeneralRepo interface {
	FindShopGeneral(
		ctx context.Context) (*shopgeneralmodel.ShopGeneral, error)
}

type printInvoiceBiz struct {
	invoiceRepo SeeInvoiceDetailRepo
	shopRepo    SeeShopGeneralRepo
	requester   middleware.Requester
}

func NewPrintInvoiceBiz(
	invoiceRepo SeeInvoiceDetailRepo,
	shopRepo SeeShopGeneralRepo,
	requester middleware.Requester) *printInvoiceBiz {
	return &printInvoiceBiz{
		invoiceRepo: invoiceRepo,
		shopRepo:    shopRepo,
		requester:   requester,
	}
}

func (biz *printInvoiceBiz) PrintInvoice(
	ctx context.Context,
	invoiceId string) (*invoicemodel.InvoicePrint, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicemodel.ErrInvoiceViewNoPermission
	}

	invoice, errInvoice := biz.invoiceRepo.SeeInvoiceDetail(
		ctx, invoiceId)
	if errInvoice != nil {
		return nil, errInvoice
	}

	general, errGetGeneral := biz.shopRepo.FindShopGeneral(ctx)
	if errGetGeneral != nil {
		return nil, errGetGeneral
	}

	invoicePrint := invoicemodel.InvoicePrint{
		Invoice:     *invoice,
		ShopGeneral: *general,
	}
	return &invoicePrint, nil
}
