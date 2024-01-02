package inventorychecknoterepo

import (
	"backend/module/ingredient/ingredientmodel"
	"backend/module/inventorychecknote/inventorychecknotemodel"
	"backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
)

type CreateInventoryCheckNoteStore interface {
	CreateInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.InventoryCheckNoteCreate,
	) error
}

type CreateInventoryCheckNoteDetailStore interface {
	CreateListInventoryCheckNoteDetail(
		ctx context.Context,
		data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate,
	) error
}

type UpdateIngredientStore interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount) error
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*ingredientmodel.Ingredient, error)
}

type StockChangeHistoryStore interface {
	CreateLisStockChangeHistory(
		ctx context.Context,
		data []stockchangehistorymodel.StockChangeHistory) error
}

type createInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       CreateInventoryCheckNoteStore
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
	ingredientStore               UpdateIngredientStore
	stockChangeHistoryStore       StockChangeHistoryStore
}

func NewCreateInventoryCheckNoteRepo(
	inventoryCheckNoteStore CreateInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore,
	ingredientStore UpdateIngredientStore,
	stockChangeHistoryStore StockChangeHistoryStore) *createInventoryCheckNoteRepo {
	return &createInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
		ingredientStore:               ingredientStore,
		stockChangeHistoryStore:       stockChangeHistoryStore,
	}
}

func (repo *createInventoryCheckNoteRepo) HandleInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	if err := repo.inventoryCheckNoteStore.CreateInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	if err := repo.inventoryCheckNoteDetailStore.CreateListInventoryCheckNoteDetail(
		ctx, data.Details,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInventoryCheckNoteRepo) HandleIngredientAmount(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	var history []stockchangehistorymodel.StockChangeHistory
	for i, value := range data.Details {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": value.IngredientId})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		data.Details[i].Initial = ingredient.Amount
		data.Details[i].Final = ingredient.Amount + value.Difference

		if data.Details[i].Final < 0 {
			return inventorychecknotemodel.ErrInventoryCheckNoteModifyAmountIsInvalid
		}

		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: value.Difference}
		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, value.IngredientId, &ingredientUpdate,
		); err != nil {
			return err
		}

		typeChange := stockchangehistorymodel.Modify
		stockChangeHistory := stockchangehistorymodel.StockChangeHistory{
			Id:           *data.Id,
			IngredientId: data.Details[i].IngredientId,
			Amount:       value.Difference,
			AmountLeft:   value.Difference + ingredient.Amount,
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
