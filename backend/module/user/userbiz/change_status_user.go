package userbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/user/usermodel"
	"context"
)

type ChangeStatusUserRepo interface {
	UpdateStatusUser(
		ctx context.Context,
		data *usermodel.UserUpdateStatus,
	) error
}

type changeStatusUserBiz struct {
	repo      ChangeStatusUserRepo
	requester middleware.Requester
}

func NewChangeStatusUserBiz(
	repo ChangeStatusUserRepo,
	requester middleware.Requester) *changeStatusUserBiz {
	return &changeStatusUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusUserBiz) ChangeStatusUser(
	ctx context.Context,
	data usermodel.UsersUpdateStatus) error {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return usermodel.ErrUserUpdateStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	for _, v := range data.UserIds {
		userUpdateStatus := usermodel.UserUpdateStatus{
			UserId:   v,
			IsActive: data.IsActive,
		}
		if err := biz.repo.UpdateStatusUser(ctx, &userUpdateStatus); err != nil {
			return err
		}
	}

	return nil
}
