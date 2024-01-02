package exportnoterepo

import (
	"backend/module/exportnote/exportnotemodel"
	"backend/module/exportnotedetail/exportnotedetailmodel"
	"backend/module/ingredient/ingredientmodel"
	"backend/module/stockchangehistory/stockchangehistorymodel"
	"backend/module/unittype/unittypemodel"
	"context"
)

type CreateExportNoteStore interface {
	CreateExportNote(
		ctx context.Context,
		data *exportnotemodel.ExportNoteCreate,
	) error
}

type CreateExportNoteDetailStore interface {
	CreateListExportNoteDetail(
		ctx context.Context,
		data []exportnotedetailmodel.ExportNoteDetailCreate,
	) error
}

type UpdateIngredientStore interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
}
type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type ListUnitTypeStore interface {
	ListUnitType(
		ctx context.Context,
		condition map[string]interface{}) ([]unittypemodel.UnitType, error)
}

type createExportNoteRepo struct {
	exportNoteStore         CreateExportNoteStore
	exportNoteDetailStore   CreateExportNoteDetailStore
	ingredientStore         UpdateIngredientStore
	stockChangeHistoryStore StockChangeHistoryStore
	unitTypeStore           ListUnitTypeStore
}

func NewCreateExportNoteRepo(
	exportNoteStore CreateExportNoteStore,
	exportNoteDetailStore CreateExportNoteDetailStore,
	ingredientStore UpdateIngredientStore,
	stockChangeHistoryStore StockChangeHistoryStore,
	unitTypeStore ListUnitTypeStore) *createExportNoteRepo {
	return &createExportNoteRepo{
		exportNoteStore:         exportNoteStore,
		exportNoteDetailStore:   exportNoteDetailStore,
		ingredientStore:         ingredientStore,
		stockChangeHistoryStore: stockChangeHistoryStore,
		unitTypeStore:           unitTypeStore,
	}
}

func (repo *createExportNoteRepo) HandleExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	if err := repo.exportNoteStore.CreateExportNote(ctx, data); err != nil {
		return err
	}

	if err := repo.exportNoteDetailStore.CreateListExportNoteDetail(
		ctx, data.ExportNoteDetails,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createExportNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	exportNoteId string,
	data *exportnotemodel.ExportNoteCreate) error {

	var history []stockchangehistorymodel.StockChangeHistory

	for _, v := range data.ExportNoteDetails {
		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: -v.AmountExportByDefaultUnitType}

		amountLeft := v.Ingredient.Amount - v.AmountExportByDefaultUnitType
		if amountLeft < 0 {
			return exportnotemodel.ErrExportNoteAmountExportIsOverTheStock
		}

		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, v.IngredientId, &ingredientUpdate,
		); err != nil {
			return err
		}

		typeChange := stockchangehistorymodel.Export
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           exportNoteId,
			IngredientId: v.IngredientId,
			Amount:       -v.AmountExportByDefaultUnitType,
			AmountLeft:   amountLeft,
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

func (repo *createExportNoteRepo) ChangeUnitOfIngredient(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	units, errListUnit := repo.unitTypeStore.ListUnitType(ctx, map[string]interface{}{})
	if errListUnit != nil {
		return errListUnit
	}

	var mapUnit map[string]unittypemodel.UnitType
	for _, v := range units {
		mapUnit[v.Id] = v
	}

	for i, v := range data.ExportNoteDetails {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": v.IngredientId})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		if mapUnit[v.UnitTypeId].MeasureType != mapUnit[ingredient.UnitTypeId].MeasureType {
			return exportnotemodel.ErrExportNoteMeasureTypeIsNotCorrect
		}

		data.ExportNoteDetails[i].Ingredient = ingredient
		data.ExportNoteDetails[i].AmountExportByDefaultUnitType =
			v.AmountExport * float32(mapUnit[ingredient.UnitTypeId].Value) / float32(mapUnit[v.UnitTypeId].Value)
		data.ExportNoteDetails[i].UnitTypeName = mapUnit[v.UnitTypeId].Name
	}
	return nil
}
