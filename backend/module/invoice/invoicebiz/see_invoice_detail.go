package invoicebiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/invoice/invoicemodel"
	"context"
)

type SeeInvoiceDetailRepo interface {
	SeeInvoiceDetail(
		ctx context.Context,
		invoiceId string) (*invoicemodel.Invoice, error)
}

type seeInvoiceDetailBiz struct {
	repo      SeeInvoiceDetailRepo
	requester middleware.Requester
}

func NewSeeInvoiceDetailBiz(
	repo SeeInvoiceDetailRepo,
	requester middleware.Requester) *seeInvoiceDetailBiz {
	return &seeInvoiceDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeInvoiceDetailBiz) SeeInvoiceDetail(
	ctx context.Context,
	invoiceId string) (*invoicemodel.Invoice, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicemodel.ErrInvoiceViewNoPermission
	}

	invoice, errInvoice := biz.repo.SeeInvoiceDetail(
		ctx, invoiceId)
	if errInvoice != nil {
		return nil, errInvoice
	}

	return invoice, nil
}
