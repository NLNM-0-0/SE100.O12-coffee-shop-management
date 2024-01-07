package exportnoterepo

import (
	"backend/common"
	"backend/module/exportnote/exportnotemodel"
	"context"
)

type ListExportNoteStore interface {
	ListExportNote(
		ctx context.Context,
		filter *exportnotemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string,
	) ([]exportnotemodel.ExportNote, error)
}

type listExportNoteRepo struct {
	store ListExportNoteStore
}

func NewListExportNoteRepo(store ListExportNoteStore) *listExportNoteRepo {
	return &listExportNoteRepo{store: store}
}

func (repo *listExportNoteRepo) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	paging *common.Paging) ([]exportnotemodel.ExportNote, error) {
	result, err := repo.store.ListExportNote(
		ctx,
		filter,
		[]string{"ExportNote.id"},
		paging,
		"CreatedByUser")

	if err != nil {
		return nil, err
	}

	return result, nil
}
