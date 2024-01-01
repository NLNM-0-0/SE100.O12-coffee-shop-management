package supplierbiz

import (
	"backend/common"
	"backend/common/enum"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStoreRepo interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string,
	) (*int, error)
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
	UpdateDebtSupplier(
		ctx context.Context,
		supplierId string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
	GetAllImportNoteId(
		ctx context.Context) ([]string, error)
}

type paySupplierBiz struct {
	gen       generator.IdGenerator
	repo      PaySupplierStoreRepo
	requester middleware.Requester
}

func NewUpdatePayBiz(
	gen generator.IdGenerator,
	repo PaySupplierStoreRepo,
	requester middleware.Requester) *paySupplierBiz {
	return &paySupplierBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *paySupplierBiz) PaySupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateDebt) (*string, error) {
	if !biz.requester.IsHasFeature(common.SupplierPayFeatureCode) {
		return nil, suppliermodel.ErrSupplierPayNoPermission
	}

	data.CreatedBy = biz.requester.GetUserId()

	if err := validateSupplierUpdateDebt(data); err != nil {
		return nil, err
	}

	debtCurrent, errGetDebt := biz.repo.GetDebtSupplier(ctx, supplierId)
	if errGetDebt != nil {
		return nil, errGetDebt
	}
	if *debtCurrent-*data.Amount < 0 {
		return nil, suppliermodel.ErrSupplierInitDebtInvalid
	}

	importNoteIds, errGetImportNoteIds := biz.repo.GetAllImportNoteId(ctx)
	if errGetImportNoteIds != nil {
		return nil, errGetImportNoteIds
	}

	supplierDebtCreate, errGetSupplierDebtCreate := getSupplierDebtCreate(
		biz.gen, supplierId, *debtCurrent, data, importNoteIds,
	)
	if errGetSupplierDebtCreate != nil {
		return nil, errGetSupplierDebtCreate
	}

	if err := biz.repo.UpdateDebtSupplier(ctx, supplierId, data); err != nil {
		return nil, err
	}

	if err := biz.repo.CreateSupplierDebt(ctx, supplierDebtCreate); err != nil {
		return nil, err
	}

	return &supplierDebtCreate.Id, nil
}

func validateSupplierUpdateDebt(data *suppliermodel.SupplierUpdateDebt) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if *data.Amount <= 0 {
		return suppliermodel.ErrSupplierDebtPayIsInvalid
	}

	return nil
}

func stringInSlice(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getSupplierDebtCreate(
	gen generator.IdGenerator,
	supplierId string,
	currentDebt int,
	data *suppliermodel.SupplierUpdateDebt,
	importNoteIds []string,
) (*supplierdebtmodel.SupplierDebtCreate, error) {
	amountPay := -*data.Amount
	amountLeft := currentDebt + amountPay

	if data.Id != nil {
		id, errGenerateId := gen.IdProcess(data.Id)
		if errGenerateId != nil {
			return nil, errGenerateId
		}
		if stringInSlice(importNoteIds, *id) {
			return nil, suppliermodel.ErrSupplierDebtIdExistedInImportNote
		}
		data.Id = id
	} else {
		id, errGenerateId := gen.GenerateId()
		if errGenerateId != nil {
			return nil, errGenerateId
		}
		for stringInSlice(importNoteIds, id) {
			id, errGenerateId = gen.GenerateId()
			if errGenerateId != nil {
				return nil, errGenerateId
			}
		}

		data.Id = &id
	}

	debtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         *data.Id,
		SupplierId: supplierId,
		Amount:     amountPay,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreatedBy:  data.CreatedBy,
	}

	return &supplierDebtCreate, nil
}
