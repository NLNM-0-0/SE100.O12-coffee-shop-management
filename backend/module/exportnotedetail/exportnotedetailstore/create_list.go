package exportnotedetailstore

import (
	"backend/common"
	"backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListExportNoteDetail(
	ctx context.Context,
	data []exportnotedetailmodel.ExportNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
