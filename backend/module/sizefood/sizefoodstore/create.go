package sizefoodstore

import (
	"backend/common"
	"backend/module/sizefood/sizefoodmodel"
	"context"
)

func (s *sqlStore) CreateSizeFood(
	ctx context.Context,
	data *sizefoodmodel.SizeFoodCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
