package importnoterepo

import (
	"backend/module/importnote/importnotemodel"
	"backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type SeeDetailImportNoteStore interface {
	ListImportNoteDetail(
		ctx context.Context,
		importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type FindImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*importnotemodel.ImportNote, error)
}

type seeImportNoteDetailRepo struct {
	importNoteStore       FindImportNoteStore
	importNoteDetailStore SeeDetailImportNoteStore
}

func NewSeeImportNoteDetailRepo(
	importNoteStore FindImportNoteStore,
	importNoteDetailStore SeeDetailImportNoteStore) *seeImportNoteDetailRepo {
	return &seeImportNoteDetailRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
	}
}

func (repo *seeImportNoteDetailRepo) SeeImportNoteDetail(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, errImportNote := repo.importNoteStore.FindImportNote(
		ctx, map[string]interface{}{"id": importNoteId},
		"Supplier", "CreatedByUser", "ClosedByUser")
	if errImportNote != nil {
		return nil, errImportNote
	}

	details, errImportNoteDetail := repo.importNoteDetailStore.ListImportNoteDetail(
		ctx,
		importNoteId,
	)
	if errImportNoteDetail != nil {
		return nil, errImportNoteDetail
	}

	importNote.Details = details

	return importNote, nil
}
