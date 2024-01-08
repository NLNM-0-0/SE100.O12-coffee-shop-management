package salereportrepo

import (
	"backend/module/invoice/invoicemodel"
	"backend/module/invoicedetail/invoicedetailmodel"
	"backend/module/product/productmodel"
	"backend/module/salereport/salereportmodel"
	"backend/module/salereportdetail/salereportetailmodel"
	"context"
	"time"
)

type FindInvoiceForReportStore interface {
	ListAllInvoiceForReport(
		ctx context.Context,
		startTime time.Time,
		endTime time.Time,
		moreKeys ...string) ([]invoicemodel.Invoice, error)
}

type ListInvoiceDetailForReportStore interface {
	ListInvoiceDetail(
		ctx context.Context,
		invoiceId string) ([]invoicedetailmodel.InvoiceDetail, error)
}

type findSaleReportRepo struct {
	invoiceStore       FindInvoiceForReportStore
	invoiceDetailStore ListInvoiceDetailForReportStore
}

func NewFindSaleReportRepo(
	invoiceStore FindInvoiceForReportStore,
	invoiceDetailStore ListInvoiceDetailForReportStore) *findSaleReportRepo {
	return &findSaleReportRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
	}
}

func (repo *findSaleReportRepo) FindSaleReport(
	ctx context.Context,
	data *salereportmodel.ReqFindSaleReport) (*salereportmodel.SaleReport, error) {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)
	data.TimeFromTime = timeFrom
	data.TimeToTime = timeTo

	allInvoices, errInvoices := repo.invoiceStore.ListAllInvoiceForReport(
		ctx, timeFrom, timeTo, "Details.Food")
	if errInvoices != nil {
		return nil, errInvoices
	}

	total := 0
	totalAmount := 0
	mapFoodAmount := make(map[string]int)
	mapFoodName := make(map[string]string)
	mapFoodSales := make(map[string]int)
	for _, invoice := range allInvoices {
		details, err := repo.invoiceDetailStore.ListInvoiceDetail(
			ctx, invoice.Id)
		if err != nil {
			return nil, err
		}

		for _, detail := range details {
			mapFoodAmount[detail.FoodId] += detail.Amount
			mapFoodName[detail.FoodId] = detail.Food.Name
			totalInvoiceDetail := detail.UnitPrice * detail.Amount
			mapFoodSales[detail.FoodId] += totalInvoiceDetail

			total += totalInvoiceDetail
			totalAmount += detail.Amount
		}
	}

	details := make([]salereportetailmodel.SaleReportDetail, 0)
	for key, value := range mapFoodName {
		if mapFoodAmount[key] != 0 {
			detail := salereportetailmodel.SaleReportDetail{
				Food: productmodel.SimpleFood{
					Id:   key,
					Name: value,
				},
				Amount:     mapFoodAmount[key],
				TotalSales: mapFoodSales[key],
			}
			details = append(details, detail)
		}
	}

	saleReport := salereportmodel.SaleReport{
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		Total:    total,
		Amount:   totalAmount,
		Details:  details,
	}

	return &saleReport, nil
}
