package customerbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/customer/customermodel"
	"context"
)

type SeeCustomerRepo interface {
	SeeCustomerDetail(
		ctx context.Context,
		customerId string) (*customermodel.Customer, error)
}

type seeCustomerDetailBiz struct {
	repo      SeeCustomerRepo
	requester middleware.Requester
}

func NewSeeCustomerDetailBiz(
	repo SeeCustomerRepo,
	requester middleware.Requester) *seeCustomerDetailBiz {
	return &seeCustomerDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeCustomerDetailBiz) SeeCustomerDetail(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	if !biz.requester.IsHasFeature(common.CustomerViewFeatureCode) {
		return nil, customermodel.ErrCustomerViewNoPermission
	}

	customer, err := biz.repo.SeeCustomerDetail(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
