package userstore

import (
	"backend/common"
	"backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateRoleUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateRole) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
