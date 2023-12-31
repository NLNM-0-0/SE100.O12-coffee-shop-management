package importnotebiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/importnote/importnotemodel"
	"context"
)

type SeeImportNoteDetailRepo interface {
	SeeImportNoteDetail(
		ctx context.Context,
		importNoteId string,
	) (*importnotemodel.ImportNote, error)
}

type seeImportNoteDetailBiz struct {
	repo      SeeImportNoteDetailRepo
	requester middleware.Requester
}

func NewSeeImportNoteDetailBiz(
	repo SeeImportNoteDetailRepo,
	requester middleware.Requester) *seeImportNoteDetailBiz {
	return &seeImportNoteDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeImportNoteDetailBiz) SeeImportNoteDetail(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	importNote, errImportNote := biz.repo.SeeImportNoteDetail(
		ctx, importNoteId)
	if errImportNote != nil {
		return nil, errImportNote
	}

	return importNote, nil
}
