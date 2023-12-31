package rolestore

import (
	"backend/common"
	"backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) UpdateRole(
	ctx context.Context,
	id string,
	data *rolemodel.RoleUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return rolemodel.ErrRoleNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
