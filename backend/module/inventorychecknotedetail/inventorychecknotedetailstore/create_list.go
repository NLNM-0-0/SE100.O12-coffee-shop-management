package inventorychecknotedetailstore

import (
	"backend/common"
	"backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListInventoryCheckNoteDetail(
	ctx context.Context,
	data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
