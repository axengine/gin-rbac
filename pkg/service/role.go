package service

import (
	"context"
	"fmt"
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
	rmas, err := svc.d.FindAllRoleMenuAction(ctx, in.RoleId)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	type rmaHit struct {
		rma *model.RoleMenuAction
		hit bool
	}
	rmaMap := make(map[string]rmaHit)
	for _, v := range rmas {
		rmaMap[fmt.Sprintf("mid_%d_aid_%d", v.MenuId, v.ActionId)] = rmaHit{
			rma: v,
			hit: false,
		}
	}
	add := make([]*model.RoleMenuAction, 0)
	for _, v := range in.MenuActions {
		k := fmt.Sprintf("mid_%d_aid_%d", v.MenuId, v.ActionId)
		rma, ok := rmaMap[k]
		if ok {
			rma.hit = true
			rmaMap[k] = rma
		} else {
			add = append(add, &model.RoleMenuAction{
				RoleId:   in.RoleId,
				MenuId:   v.MenuId,
				ActionId: v.ActionId,
			})
		}
	}
	del := make([]int64, 0)
	for _, v := range rmaMap {
		if !v.hit {
			del = append(del, v.rma.Id)
		}
	}

	if len(add) > 0 || len(del) > 0 {
		if err := svc.d.UpdateRoleMenuAction(ctx, add, del); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
	}

	return nil
}
