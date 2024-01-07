package importnoterepo

import (
	"backend/common/enum"
	"backend/module/importnote/importnotemodel"
	"backend/module/importnotedetail/importnotedetailmodel"
	"backend/module/ingredient/ingredientmodel"
	"backend/module/stockchangehistory/stockchangehistorymodel"
	"backend/module/supplier/suppliermodel"
	"backend/module/supplierdebt/supplierdebtmodel"
	"backend/module/unittype/unittypemodel"
	"context"
)

type ChangeStatusImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		id string,
		data *importnotemodel.ImportNoteUpdate,
	) error
}

type GetImportNoteDetailStore interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateAmountIngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*ingredientmodel.Ingredient, error)
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
}

type UpdateDebtOfSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}

type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore         ChangeStatusImportNoteStore
	importNoteDetailStore   GetImportNoteDetailStore
	ingredientStore         UpdateAmountIngredientStore
	supplierStore           UpdateDebtOfSupplierStore
	supplierDebtStore       CreateSupplierDebtStore
	stockChangeHistoryStore StockChangeHistoryStore
	unitTypeStore           ListUnitTypeStore
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStore,
	importNoteDetailStore GetImportNoteDetailStore,
	ingredientStore UpdateAmountIngredientStore,
	supplierStore UpdateDebtOfSupplierStore,
	supplierDebtStore CreateSupplierDebtStore,
	stockChangeHistoryStore StockChangeHistoryStore,
	unitTypeStore ListUnitTypeStore) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:         importNoteStore,
		importNoteDetailStore:   importNoteDetailStore,
		ingredientStore:         ingredientStore,
		supplierStore:           supplierStore,
		supplierDebtStore:       supplierDebtStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
		unitTypeStore:           unitTypeStore,
	}
}

func (repo *changeStatusImportNoteRepo) FindImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, err := repo.importNoteStore.FindImportNote(ctx, map[string]interface{}{"id": importNoteId})
	if err != nil {
		return nil, err
	}
	return importNote, nil
}

func (repo *changeStatusImportNoteRepo) UpdateImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	if err := repo.importNoteStore.UpdateImportNote(ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ImportNoteUpdate) error {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx,
		map[string]interface{}{"id": importNote.SupplierId})
	if err != nil {
		return err
	}

	amountBorrow := importNote.TotalPrice
	amountLeft := supplier.Debt + amountBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         supplierDebtId,
		SupplierId: importNote.SupplierId,
		Amount:     amountBorrow,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreatedBy:  importNote.ClosedBy,
	}

	if err := repo.supplierDebtStore.CreateSupplierDebt(
		ctx, &supplierDebtCreate,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) UpdateDebtSupplier(
	ctx context.Context,
	importNote *importnotemodel.ImportNoteUpdate) error {
	amount := importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Id:        &importNote.Id,
		Amount:    &amount,
		CreatedBy: importNote.ClosedBy,
	}
	if err := repo.supplierStore.UpdateSupplierDebt(
		ctx, importNote.SupplierId, &supplierUpdateDebt,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) FindListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	importNoteDetails, errGetImportNoteDetails := repo.importNoteDetailStore.FindListImportNoteDetail(
		ctx,
		map[string]interface{}{"importNoteId": importNoteId})
	if errGetImportNoteDetails != nil {
		return nil, errGetImportNoteDetails
	}
	return importNoteDetails, nil
}

func (repo *changeStatusImportNoteRepo) HandleIngredient(
	ctx context.Context,
	importNoteId string,
	details []importnotedetailmodel.ImportNoteDetail) error {
	var history []stockchangehistorymodel.StockChangeHistory
	for _, v := range details {
		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: v.AmountImportByDefaultUnitType}

		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, v.IngredientId, &ingredientUpdate,
		); err != nil {
			return err
		}

		typeChange := stockchangehistorymodel.Import
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           importNoteId,
			IngredientId: v.IngredientId,
			Amount:       v.AmountImportByDefaultUnitType,
			AmountLeft:   v.AmountImportByDefaultUnitType + v.TempIngredient.Amount,
			Type:         &typeChange,
		}
		history = append(history, stockChangeHistory)
	}

	if err := repo.stockChangeHistoryStore.CreateLisStockChangeHistory(
		ctx, history); err != nil {
		return err
	}

	return nil
}

func (repo *changeStatusImportNoteRepo) ChangeUnitOfIngredient(
	ctx context.Context,
	details []importnotedetailmodel.ImportNoteDetail) error {
	units, errListUnit := repo.unitTypeStore.ListUnitType(ctx, map[string]interface{}{})
	if errListUnit != nil {
		return errListUnit
	}

	mapUnit := make(map[string]unittypemodel.UnitType)
	mapUnitName := make(map[string]unittypemodel.UnitType)
	for _, v := range units {
		mapUnit[v.Id] = v
		mapUnitName[v.Name] = v
	}

	for i, v := range details {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": v.IngredientId})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		if mapUnitName[v.UnitTypeName].MeasureType != mapUnit[ingredient.UnitTypeId].MeasureType {
			return importnotemodel.ErrImportNoteMeasureTypeIsNotCorrect
		}

		details[i].TempIngredient = ingredient
		details[i].AmountImportByDefaultUnitType =
			v.AmountImport *
				float32(mapUnit[ingredient.UnitTypeId].Value) /
				float32(mapUnitName[v.UnitTypeName].Value)
	}
	return nil
}
