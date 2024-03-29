package invoicebiz

import (
	"backend/common"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/customer/customermodel"
	"backend/module/invoice/invoicemodel"
	"backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"fmt"
)

type CreateInvoiceRepo interface {
	GetShopGeneral(
		ctx context.Context,
	) (*shopgeneralmodel.ShopGeneral, error)
	HandleCheckPermissionStatus(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleData(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	FindCustomer(
		ctx context.Context,
		customerId string,
	) (*customermodel.Customer, error)
	UpdateCustomerPoint(
		ctx context.Context,
		customerId string,
		data customermodel.CustomerUpdatePoint,
	) error
	HandleInvoice(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		invoiceId string,
		ingredientTotalAmountNeedUpdate map[string]float32,
	) error
}

type createInvoiceBiz struct {
	gen       generator.IdGenerator
	repo      CreateInvoiceRepo
	requester middleware.Requester
}

func NewCreateInvoiceBiz(
	gen generator.IdGenerator,
	repo CreateInvoiceRepo,
	requester middleware.Requester) *createInvoiceBiz {
	return &createInvoiceBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInvoiceBiz) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if !biz.requester.IsHasFeature(common.InvoiceCreateFeatureCode) {
		return invoicemodel.ErrInvoiceCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInvoiceId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleCheckPermissionStatus(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleData(ctx, data); err != nil {
		return err
	}

	general, errGetShopGeneral := biz.repo.GetShopGeneral(ctx)
	if errGetShopGeneral != nil {
		return errGetShopGeneral
	}

	if err := biz.repo.HandleIngredientTotalAmount(
		ctx, data.Id, data.MapIngredient); err != nil {
		return err
	}

	if data.CustomerId != "" {
		customer, errGetCustomer := biz.repo.FindCustomer(ctx, data.CustomerId)
		if errGetCustomer != nil {
			return errGetCustomer
		}

		priceUseForPoint := float32(0)
		pointUse := 0
		if data.IsUsePoint {
			if float32(data.TotalPrice) >= float32(customer.Point)*general.UsePointPercent {
				pointUse = customer.Point
				priceUseForPoint = float32(customer.Point) * general.UsePointPercent
			} else {
				pointUse = common.RoundToInt(float32(data.TotalPrice) / general.UsePointPercent)
				priceUseForPoint = float32(data.TotalPrice)
			}
		}

		priceUseForPointInt := common.RoundToInt(priceUseForPoint)
		common.CustomRound(&priceUseForPoint)
		data.AmountReceived = data.TotalPrice - priceUseForPointInt
		data.AmountPriceUsePoint = priceUseForPointInt

		pointReceive := common.RoundToInt(float32(data.AmountReceived) * general.AccumulatePointPercent)
		fmt.Println(pointReceive)

		data.PointUse = pointUse
		data.PointReceive = pointReceive

		amountPointNeedUpdate :=
			pointReceive - pointUse

		customerUpdatePoint := customermodel.CustomerUpdatePoint{
			Amount: &amountPointNeedUpdate,
		}
		if err := biz.repo.UpdateCustomerPoint(
			ctx, data.CustomerId, customerUpdatePoint); err != nil {
			return err
		}

	} else {
		data.AmountReceived = data.TotalPrice
		if data.IsUsePoint {
			return invoicemodel.ErrInvoiceNotHaveCustomerToUsePoint
		}
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return err
	}

	fmt.Println(data.TotalCost)
	return nil
}

func handleInvoiceId(gen generator.IdGenerator, data *invoicemodel.InvoiceCreate) error {
	idInvoice, err := gen.GenerateId()
	if err != nil {
		return err
	}
	data.Id = idInvoice

	for i := range data.InvoiceDetails {
		data.InvoiceDetails[i].InvoiceId = idInvoice
	}

	return err
}
