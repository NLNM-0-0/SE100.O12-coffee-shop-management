package inventorychecknotedetailstore

import (
	"backend/common"
	"backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListInventoryCheckNoteDetail(
	ctx context.Context,
	inventoryCheckNoteId string) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error) {
	var result []inventorychecknotedetailmodel.InventoryCheckNoteDetail
	db := s.db

	db = db.Table(common.TableInventoryCheckNoteDetail)

	db = db.Where("inventoryCheckNoteId = ?", inventoryCheckNoteId)

	if err := db.
		Preload("Ingredient.UnitType", func(db *gorm.DB) *gorm.DB {
			return db.Order("Ingredient.name")
		}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
