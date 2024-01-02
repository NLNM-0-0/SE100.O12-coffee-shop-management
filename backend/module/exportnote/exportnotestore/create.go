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
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return exportnotemodel.ErrExportNoteIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
