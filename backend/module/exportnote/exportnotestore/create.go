package exportnotestore

import (
	"backend/common"
	"backend/module/exportnote/exportnotemodel"
	"context"
)

func (s *sqlStore) CreateExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
