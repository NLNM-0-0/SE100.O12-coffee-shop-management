package stockchangehistorystore

import (
	"backend/common"
	"backend/module/stockchangehistory/stockchangehistorymodel"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (s *sqlStore) GetNearlyStockChangeHistory(
	ctx context.Context,
	ingredientId string,
	timeFrom time.Time) (*stockchangehistorymodel.StockChangeHistory, error) {
	var result stockchangehistorymodel.StockChangeHistory
	db := s.db

	db = db.Table(common.TableStockChangeHistory)

	timeRequest := timeFrom.Add(-time.Second)

	if err := db.
		Where("ingredientId = ?", ingredientId).
		Where("createdAt <= ?", timeRequest).
		Order("createdAt desc").
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
