package userbiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/user/usermodel"
	"context"
)

type UpdateInfoUserRepo interface {
	CheckUserStatusPermission(
		ctx context.Context,
		userId string,
	) error
	UpdateInfoUser(
		ctx context.Context,
		userId string,
		data *usermodel.UserUpdateInfo,
	) error
}

type updateInfoUserBiz struct {
	repo      UpdateInfoUserRepo
	requester middleware.Requester
}

func NewUpdateInfoUserBiz(
	repo UpdateInfoUserRepo,
	requester middleware.Requester) *updateInfoUserBiz {
	return &updateInfoUserBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *updateInfoUserBiz) UpdateUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.UserUpdateInfoFeatureCode) {
		return usermodel.ErrUserUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckUserStatusPermission(ctx, id); err != nil {
		return err
	}

	if err := biz.repo.UpdateInfoUser(ctx, id, data); err != nil {
		return err
	}

	return nil
}
