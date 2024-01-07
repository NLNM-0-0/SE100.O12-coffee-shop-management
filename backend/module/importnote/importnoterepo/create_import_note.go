package importnoterepo

import (
	"backend/common"
	"backend/module/importnote/importnotemodel"
	"backend/module/importnotedetail/importnotedetailmodel"
	"backend/module/ingredient/ingredientmodel"
	"backend/module/unittype/unittypemodel"
	"context"
)

type CreateImportNoteStore interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
}

type CreateImportNoteDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []importnotedetailmodel.ImportNoteDetailCreate,
	) error
}

type UpdatePriceIngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
	UpdatePriceIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdatePrice,
	) error
}

type ListUnitTypeStore interface {
	ListUnitType(
		ctx context.Context,
		condition map[string]interface{}) ([]unittypemodel.UnitType, error)
}

type createImportNoteRepo struct {
	importNoteStore       CreateImportNoteStore
	importNoteDetailStore CreateImportNoteDetailStore
	ingredientStore       UpdatePriceIngredientStore
	unitTypeStore         ListUnitTypeStore
}

func NewCreateImportNoteRepo(
	importNoteStore CreateImportNoteStore,
	importNoteDetailStore CreateImportNoteDetailStore,
	updatePriceIngredientStore UpdatePriceIngredientStore,
	unitTypeStore ListUnitTypeStore) *createImportNoteRepo {
	return &createImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		ingredientStore:       updatePriceIngredientStore,
		unitTypeStore:         unitTypeStore,
	}
}

func (repo *createImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if err := repo.importNoteStore.CreateImportNote(ctx, data); err != nil {
		return err
	}
	if err := repo.importNoteDetailStore.CreateListImportNoteDetail(
		ctx,
		data.ImportNoteDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) UpdatePriceIngredient(
	ctx context.Context,
	ingredientId string,
	price int) error {
	ingredientUpdatePrice := ingredientmodel.IngredientUpdatePrice{
		Price: &price,
	}

	if err := repo.ingredientStore.UpdatePriceIngredient(
		ctx, ingredientId, &ingredientUpdatePrice,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) ChangeUnitOfIngredient(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	units, errListUnit := repo.unitTypeStore.ListUnitType(ctx, map[string]interface{}{})
	if errListUnit != nil {
		return errListUnit
	}

	mapUnit := make(map[string]unittypemodel.UnitType)
	for _, v := range units {
		mapUnit[v.Id] = v
	}

	for i, v := range data.ImportNoteDetails {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": v.IngredientId})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		if mapUnit[v.UnitTypeId].MeasureType != mapUnit[ingredient.UnitTypeId].MeasureType {
			return importnotemodel.ErrImportNoteMeasureTypeIsNotCorrect
		}

		data.ImportNoteDetails[i].Ingredient = ingredient
		data.ImportNoteDetails[i].UnitTypeName = mapUnit[v.UnitTypeId].Name
		data.ImportNoteDetails[i].PriceByDefaultUnitType =
			common.RoundToInt(float32(v.Price) * float32(mapUnit[v.UnitTypeId].Value) / float32(mapUnit[ingredient.UnitTypeId].Value))
	}
	return nil
}
