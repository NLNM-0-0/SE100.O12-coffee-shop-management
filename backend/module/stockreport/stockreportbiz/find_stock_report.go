package stockreportbiz

import (
	"backend/common"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/ingredient/ingredientmodel"
	"backend/module/stockchangehistory/stockchangehistorymodel"
	"backend/module/stockreport/stockreportmodel"
	"backend/module/stockreportdetail/stockreportdetailmodel"
	"context"
	"errors"
	"time"
)

type FindIngredientStore interface {
	GetAllSimpleIngredient(
		ctx context.Context,
		moreKeys ...string) ([]ingredientmodel.SimpleIngredient, error)
}

type FindStockChangeHistoryStore interface {
	ListAllStockChangeForReport(
		ctx context.Context,
		ingredientId string,
		timeFrom time.Time,
		timeTo time.Time) ([]stockchangehistorymodel.StockChangeHistory, error)
	GetNearlyStockChangeHistory(
		ctx context.Context,
		ingredientId string,
		timeFrom time.Time) (*stockchangehistorymodel.StockChangeHistory, error)
}

type FindStockReportStore interface {
	FindStockReport(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*stockreportmodel.StockReport, error)
	CreateStockReport(
		ctx context.Context,
		data *stockreportmodel.ReqFindStockReport) error
}

type findStockReportBiz struct {
	gen                     generator.IdGenerator
	ingredientStore         FindIngredientStore
	stockChangeHistoryStore FindStockChangeHistoryStore
	inventoryReportStore    FindStockReportStore
	requester               middleware.Requester
}

func NewFindStockReportBiz(
	gen generator.IdGenerator,
	ingredientStore FindIngredientStore,
	stockChangeHistoryStore FindStockChangeHistoryStore,
	inventoryReportStore FindStockReportStore,
	requester middleware.Requester) *findStockReportBiz {
	return &findStockReportBiz{
		gen:                     gen,
		ingredientStore:         ingredientStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
		inventoryReportStore:    inventoryReportStore,
		requester:               requester,
	}
}

func (biz *findStockReportBiz) FindStockReport(
	ctx context.Context,
	data *stockreportmodel.ReqFindStockReport,
) (*stockreportmodel.StockReport, error) {
	if !biz.requester.IsHasFeature(common.StockReportViewFeatureCode) {
		return nil, stockreportmodel.ErrStockReportViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	timeFrom := time.Unix(data.TimeFrom, 0)
	timeTo := time.Unix(data.TimeTo, 0)
	data.TimeFromTime = timeFrom
	data.TimeToTime = timeTo

	report, err := biz.inventoryReportStore.FindStockReport(
		ctx, map[string]interface{}{
			"timeFrom": data.TimeFromTime, "timeTo": data.TimeToTime,
		}, "Details.Ingredient",
	)
	if err == nil {
		return report, nil
	} else {
		var appErr *common.AppError
		if errors.As(err, &appErr) {
			if appErr.Key != common.ErrRecordNotFound().Key {
				return nil, err
			}
		}
	}

	allIngredient, err := biz.ingredientStore.GetAllSimpleIngredient(ctx, "UnitType")
	if err != nil {
		return nil, err
	}

	reportId := ""

	now := time.Now()

	if now.Before(timeFrom) {
		return nil, stockreportmodel.ErrStockReportFutureDateInvalid
	}

	if timeTo.Before(now) {
		id, err := biz.gen.GenerateId()
		if err != nil {
			return nil, err
		}
		reportId = id
	}

	allDetailCreates := make([]stockreportdetailmodel.StockReportDetailCreate, 0)
	allDetails := make([]stockreportdetailmodel.StockReportDetail, 0)
	for _, ingredient := range allIngredient {
		stockChange, err :=
			biz.stockChangeHistoryStore.ListAllStockChangeForReport(
				ctx, ingredient.Id, timeFrom, timeTo)
		if err != nil {
			return nil, err
		}

		sellAmount := float32(0)
		importAmount := float32(0)
		exportAmount := float32(0)
		modifyAmount := float32(0)
		for _, change := range stockChange {
			if *change.Type == stockchangehistorymodel.Sell {
				sellAmount += change.Amount
			} else if *change.Type == stockchangehistorymodel.Import {
				importAmount += change.Amount
			} else if *change.Type == stockchangehistorymodel.Modify {
				modifyAmount += change.Amount
			} else if *change.Type == stockchangehistorymodel.Export {
				exportAmount += change.Amount
			}
		}

		initial := float32(0)
		if nearly, err :=
			biz.stockChangeHistoryStore.GetNearlyStockChangeHistory(
				ctx, ingredient.Id, timeFrom,
			); err != nil {
			var appErr *common.AppError
			if errors.As(err, &appErr) {
				if appErr.Key != common.ErrRecordNotFound().Key {
					return nil, err
				}
			}
		} else {
			initial = nearly.AmountLeft
		}

		final := initial
		if len(stockChange) != 0 {
			final = stockChange[len(stockChange)-1].AmountLeft
		}

		if initial == 0 {
			initial = final - importAmount - sellAmount - modifyAmount
		}

		if initial != 0 || sellAmount != 0 ||
			importAmount != 0 || modifyAmount != 0 {
			detailCreate := stockreportdetailmodel.StockReportDetailCreate{
				ReportId:     reportId,
				IngredientId: ingredient.Id,
				Initial:      initial,
				Sell:         sellAmount,
				Import:       importAmount,
				Export:       exportAmount,
				Modify:       modifyAmount,
				Final:        final,
			}
			allDetailCreates = append(allDetailCreates, detailCreate)

			detail := stockreportdetailmodel.StockReportDetail{
				ReportId:     reportId,
				IngredientId: ingredient.Id,
				Ingredient: ingredientmodel.SimpleIngredient{
					Id:         ingredient.Id,
					Name:       ingredient.Name,
					UnitTypeId: ingredient.UnitTypeId,
					UnitType:   ingredient.UnitType,
				},
				Initial: initial,
				Sell:    sellAmount,
				Import:  importAmount,
				Export:  exportAmount,
				Modify:  modifyAmount,
				Final:   final,
			}
			allDetails = append(allDetails, detail)
		}
	}

	data.Id = reportId
	data.Details = allDetailCreates
	if reportId != "" {
		if err := biz.inventoryReportStore.CreateStockReport(
			ctx, data,
		); err != nil {
			return nil, err
		}
	}

	stockReport := stockreportmodel.StockReport{
		Id:       reportId,
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		Details:  allDetails,
	}

	return &stockReport, nil
}
