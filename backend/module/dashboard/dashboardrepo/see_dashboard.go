package dashboardrepo

import (
	"backend/middleware"
	"backend/module/dashboard/dashboardmodel"
	"backend/module/invoice/invoicemodel"
	"backend/module/invoicedetail/invoicedetailmodel"
	"backend/module/product/productmodel"
	"backend/module/salereport/salereportmodel"
	"backend/module/salereportdetail/salereportetailmodel"
	"context"
	"sort"
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

type seeDashboardRepo struct {
	invoiceStore       FindInvoiceForReportStore
	invoiceDetailStore ListInvoiceDetailForReportStore
	requester          middleware.Requester
}

func NewSeeDashboardBiz(
	invoiceStore FindInvoiceForReportStore,
	invoiceDetailStore ListInvoiceDetailForReportStore) *seeDashboardRepo {
	return &seeDashboardRepo{
		invoiceStore:       invoiceStore,
		invoiceDetailStore: invoiceDetailStore,
	}
}

func (repo *seeDashboardRepo) SeeDashboard(
	ctx context.Context,
	data *dashboardmodel.ReqSeeDashboard) (*dashboardmodel.ResSeeDashboard, error) {
	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)

	resDashBoard := dashboardmodel.ResSeeDashboard{
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
	}

	allInvoices, errInvoices := repo.invoiceStore.ListAllInvoiceForReport(
		ctx, timeFrom, timeTo, "Details.Food")
	if errInvoices != nil {
		return nil, errInvoices
	}

	total := 0
	totalAmount := 0
	totalCustomer := 0
	totalPoint := 0
	mapFoodAmount := make(map[string]int)
	mapFoodName := make(map[string]string)
	mapFoodSales := make(map[string]int)
	listCost := make([]dashboardmodel.ChartComponent, 0)
	listPrice := make([]dashboardmodel.ChartComponent, 0)
	for _, invoice := range allInvoices {
		chartSale := dashboardmodel.ChartComponent{
			Time:  *invoice.CreatedAt,
			Value: invoice.TotalCost,
		}
		listCost = append(listCost, chartSale)

		chartAmountReceive := dashboardmodel.ChartComponent{
			Time:  *invoice.CreatedAt,
			Value: invoice.TotalPrice,
		}
		listPrice = append(listPrice, chartAmountReceive)

		totalPoint += invoice.PointReceive

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

		if invoice.CustomerId != "" {
			totalCustomer++
		}
	}

	details := make(salereportmodel.Details, 0)
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

	sort.Sort(details)
	listFood := make([]productmodel.FoodWithAmount, 0)
	for _, v := range details {
		food := productmodel.FoodWithAmount{
			Id:     v.Food.Id,
			Name:   v.Food.Name,
			Amount: v.TotalSales,
		}
		listFood = append(listFood, food)
	}
	resDashBoard.TopSoldFoods = listFood

	resDashBoard.ChartCostComponents = listCost
	resDashBoard.ChartPriceComponents = listPrice
	resDashBoard.TotalSale = total
	resDashBoard.TotalCustomer = totalCustomer
	resDashBoard.TotalSold = totalAmount
	resDashBoard.TotalPoint = totalPoint

	return &resDashBoard, nil
}
