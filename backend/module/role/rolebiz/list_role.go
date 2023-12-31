package rolebiz

import (
	"backend/common"
	"backend/middleware"
	"backend/module/role/rolemodel"
	"context"
)

type ListRoleRepo interface {
	ListRole(
		ctx context.Context,
	) ([]rolemodel.SimpleRole, error)
}

type listRoleBiz struct {
	repo      ListRoleRepo
	requester middleware.Requester
}

func NewListRoleBiz(
	repo ListRoleRepo,
	requester middleware.Requester) *listRoleBiz {
	return &listRoleBiz{repo: repo, requester: requester}
}

func (biz *listRoleBiz) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	if biz.requester.GetRoleId() != common.RoleAdminId {
		return nil, rolemodel.ErrRoleViewNoPermission
	}

	result, err := biz.repo.ListRole(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
