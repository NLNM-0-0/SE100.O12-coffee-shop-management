package importnotestore

import (
	"backend/common"
	"backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) UpdateImportNote(
	ctx context.Context,
	id string,
	data *importnotemodel.ImportNoteUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
