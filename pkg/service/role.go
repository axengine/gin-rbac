package service

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/errc"
)

func (svc *Service) CreateRoleConfig(ctx context.Context, in *model.CreateRoleConfigReq) error {
	r := &model.RoleConfig{
		AppId:  in.AppId,
		Name:   in.Name,
		IsRoot: in.IsRoot,
		Status: types.LimitNormal,
		Memo:   in.Memo,
	}
	if err := svc.d.CreateRoleConfig(ctx, r); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	return nil
}

func (svc *Service) UpsertRoleMenuAction(ctx context.Context, in *model.UpsertRoleMenuActionReq) error {
	return nil
}
