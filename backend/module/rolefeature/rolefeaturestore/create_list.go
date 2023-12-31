package rolefeaturestore

import (
	"backend/common"
	"backend/module/rolefeature/rolefeaturemodel"
	"context"
)

func (s *sqlStore) CreateListRoleFeatureDetail(
	ctx context.Context,
	data []rolefeaturemodel.RoleFeature) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
