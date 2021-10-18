package service

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/errc"
)

func (svc *Service) CreateMenuConfig(ctx context.Context, in *model.CreateMenuConfigReq) error {
	r := &model.MenuConfig{
		AppId:    in.AppId,
		Name:     in.Name,
		Memo:     in.Memo,
		ParentId: in.ParentId,
		Status:   types.LimitNormal,
	}
	r.Actions = new(types.IntSliceStr).Marshal(in.Actions)

	if err := svc.d.CreateMenuConfig(ctx, r); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpsertActionConfig(ctx context.Context, in *model.UpsertActionConfigReq) error {
	if err := svc.d.UpsertActionConfig(ctx, in); err != nil {
		return errc.WithStack(err)
	}
	return nil
}
