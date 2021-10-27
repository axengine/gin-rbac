package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
)

func (svc *Service) ListRoleConfig(ctx context.Context, in *model.ListRoleConfigReq, out *typ.ListResp) error {
	c, records, err := svc.d.ListRoleConfig(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	list := make([]*model.ListRoleConfig, 0, len(records))
	for _, v := range records {
		d := &model.ListRoleConfig{
			Id:        v.Id,
			AppId:     v.AppId,
			AppName:   "",
			Name:      v.Name,
			IsRoot:    v.IsRoot,
			Memo:      v.Memo,
			UpdatedAt: v.UpdatedAt.Unix(),
		}
		if exists, app, err := svc.d.GetAppConfigFromCache(ctx, &model.GetAppConfigReq{
			AppId: v.AppId,
		}); err == nil && exists {
			d.AppName = app.Name
		}
		list = append(list, d)
	}

	out.Count = c
	out.List = list
	return nil
}

func (svc *Service) GetRoleMenuAction(ctx context.Context, in *model.GetRoleMenuActionReq, out *model.GetRoleMenuActionResp) error {
	exists, role, err := svc.d.GetRoleConfig(ctx, &model.GetRoleConfigReq{
		Id: in.RoleId,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("roleId")
	}

	actions, err := svc.d.FindRoleAllMenuAction(ctx, role.Id)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	menusMap := make(map[int64][]int64)
	for _, v := range actions {
		vv, ok := menusMap[v.MenuId]
		if ok {
			menusMap[v.MenuId] = append(vv, v.ActionId)
		} else {
			menusMap[v.MenuId] = []int64{v.ActionId}
		}
	}
	for mid, aid := range menusMap {
		out.MenuActions = append(out.MenuActions, model.MenuAction{
			MenuId:  mid,
			Actions: aid,
		})
	}
	return nil
}

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

func (svc *Service) UpdateRoleConfig(ctx context.Context, in *model.UpdateRoleConfigReq) error {
	d := &model.RoleConfig{
		Id: in.Id,
	}
	cols := make([]string, 0)
	if len(in.Name) > 0 {
		cols = append(cols, "name")
		d.Name = in.Name
	}
	if len(in.Memo) > 0 {
		cols = append(cols, "memo")
		d.Memo = in.Memo
	}
	if in.Status > 0 {
		cols = append(cols, "status")
		d.Status = in.Status
	}

	if in.IsRoot > 0 {
		cols = append(cols, "is_root")
		d.IsRoot = in.IsRoot
	}

	if err := svc.d.UpdateRoleConfig(ctx, d, cols); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpsertRoleMenuAction(ctx context.Context, in *model.UpsertRoleMenuActionReq) error {
	rmas, err := svc.d.FindRoleAllMenuAction(ctx, in.RoleId)
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
		for _, aid := range v.Actions {
			k := fmt.Sprintf("mid_%d_aid_%d", v.MenuId, aid)
			rma, ok := rmaMap[k]
			if ok {
				rma.hit = true
				rmaMap[k] = rma
			} else {
				add = append(add, &model.RoleMenuAction{
					RoleId:   in.RoleId,
					MenuId:   v.MenuId,
					ActionId: aid,
				})
			}
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

func (svc *Service) DelRoleConfig(ctx context.Context) error {
	return nil
}
