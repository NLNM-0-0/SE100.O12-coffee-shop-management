package userrepo

import (
	"backend/module/user/usermodel"
	"context"
)

type ChangeStatusUserStore interface {
	UpdateStatusUser(
		ctx context.Context,
		data *usermodel.UserUpdateStatus) error
}

type changeStatusUserRepo struct {
	userStore ChangeStatusUserStore
}

func NewChangeStatusUserRepo(
	userStore ChangeStatusUserStore) *changeStatusUserRepo {
	return &changeStatusUserRepo{
		userStore: userStore,
	}
}

func (repo *changeStatusUserRepo) UpdateStatusUser(
	ctx context.Context,
	data *usermodel.UserUpdateStatus) error {
	if err := repo.userStore.UpdateStatusUser(ctx, data); err != nil {
		return err
	}
	return nil
}
