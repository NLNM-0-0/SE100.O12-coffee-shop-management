package importnotebiz

import (
	"backend/common"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/importnote/importnotemodel"
	"context"
)

type CreateImportNoteRepo interface {
	HandleCreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
	UpdatePriceIngredient(
		ctx context.Context,
		ingredientId string,
		price int,
	) error
	ChangeUnitOfIngredient(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
}

type createImportNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateImportNoteRepo
	requester middleware.Requester
}

func NewCreateImportNoteBiz(
	gen generator.IdGenerator,
	repo CreateImportNoteRepo,
	requester middleware.Requester) *createImportNoteBiz {
	return &createImportNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createImportNoteBiz) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if !biz.requester.IsHasFeature(common.ImportNoteCreateFeatureCode) {
		return importnotemodel.ErrImportNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleImportNoteCreateId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.ChangeUnitOfIngredient(ctx, data); err != nil {
		return err
	}

	handleTotalPrice(data)

	if err := biz.repo.HandleCreateImportNote(ctx, data); err != nil {
		return err
	}

	for _, v := range data.ImportNoteDetails {
		if v.IsReplacePrice {
			if err := biz.repo.UpdatePriceIngredient(
				ctx, v.IngredientId, v.PriceByDefaultUnitType,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleImportNoteCreateId(
	gen generator.IdGenerator,
	data *importnotemodel.ImportNoteCreate) error {
	idImportNote, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idImportNote

	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].ImportNoteId = *idImportNote
	}
	return nil
}

func handleTotalPrice(data *importnotemodel.ImportNoteCreate) {
	var totalPrice = 0
	for i, importNoteDetail := range data.ImportNoteDetails {
		totalUnit := float32(importNoteDetail.Price) * importNoteDetail.AmountImport
		totalUnitInt := common.RoundToInt(totalUnit)
		data.ImportNoteDetails[i].TotalUnit = totalUnitInt
		totalPrice += totalUnitInt
	}

	data.TotalPrice = totalPrice
}
