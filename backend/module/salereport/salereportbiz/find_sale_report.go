package salereportbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/salereport/salereportmodel"
	"context"
)

type FindSaleReportRepo interface {
	FindSaleReport(
		ctx context.Context,
		data *salereportmodel.ReqFindSaleReport) (*salereportmodel.SaleReport, error)
}

type findSaleReportBiz struct {
	repo      FindSaleReportRepo
	requester middleware.Requester
}

func NewFindSaleReportBiz(
	repo FindSaleReportRepo,
	requester middleware.Requester) *findSaleReportBiz {
	return &findSaleReportBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *findSaleReportBiz) FindSaleReport(
	ctx context.Context,
	data *salereportmodel.ReqFindSaleReport) (*salereportmodel.SaleReport, error) {
	if !biz.requester.IsHasFeature(common.SaleReportViewFeatureCode) {
		return nil, salereportmodel.ErrSaleReportViewNoPermission
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	saleReport, err := biz.repo.FindSaleReport(ctx, data)
	if err != nil {
		return nil, err
	}

	return saleReport, nil
}
