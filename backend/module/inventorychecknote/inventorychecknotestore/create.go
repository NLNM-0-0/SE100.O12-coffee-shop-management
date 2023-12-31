package inventorychecknotestore

import (
	"backend/common"
	"backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

func (s *sqlStore) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
