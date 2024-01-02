package exportnotebiz

import (
	"backend/common"
	"backend/component/generator"
	"backend/middleware"
	"backend/module/exportnote/exportnotemodel"
	"context"
)

type CreateExportNoteRepo interface {
	HandleExportNote(
		ctx context.Context,
		data *exportnotemodel.ExportNoteCreate,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		exportNoteId string,
		data *exportnotemodel.ExportNoteCreate,
	) error
	ChangeUnitOfIngredient(
		ctx context.Context,
		data *exportnotemodel.ExportNoteCreate,
	) error
}

type createExportNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateExportNoteRepo
	requester middleware.Requester
}

func NewCreateExportNoteBiz(
	gen generator.IdGenerator,
	repo CreateExportNoteRepo,
	requester middleware.Requester) *createExportNoteBiz {
	return &createExportNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createExportNoteBiz) CreateExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	if !biz.requester.IsHasFeature(common.ExportNoteCreateFeatureCode) {
		return exportnotemodel.ErrExportNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleExportNoteId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.ChangeUnitOfIngredient(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleExportNote(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleIngredientTotalAmount(
		ctx, *data.Id, data); err != nil {
		return err
	}

	return nil
}

func handleExportNoteId(
	gen generator.IdGenerator,
	data *exportnotemodel.ExportNoteCreate) error {
	idCancelNote, errGenerateIdCancelNote := gen.IdProcess(data.Id)
	if errGenerateIdCancelNote != nil {
		return errGenerateIdCancelNote
	}
	data.Id = idCancelNote

	for i := range data.ExportNoteDetails {
		data.ExportNoteDetails[i].ExportNoteId = *idCancelNote
	}

	return nil
}
