package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/errc"
	"sort"
	"strings"
)

func (svc *Service) CreateMenuConfig(ctx context.Context, in *model.CreateMenuConfigReq) error {
	r := &model.MenuConfig{
		AppId:    in.AppId,
		Name:     in.Name,
		Memo:     in.Memo,
		ParentId: in.ParentId,
		Status:   types.LimitNormal,
		Sequence: in.Sequence,
	}
	r.Actions = new(types.IntSplitStr).Marshal(in.Actions)

	if err := svc.d.CreateMenuConfig(ctx, r); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpsertActionConfig(ctx context.Context, in *model.UpsertActionConfigReq) error {
	in.Method = strings.ToUpper(in.Method)
	if err := svc.d.UpsertActionConfig(ctx, in); err != nil {
		return errc.WithStack(err)
	}
	return nil
}

func (svc *Service) UpdateMenuAction(ctx context.Context, in *model.UpdateMenuActionReq) error {
	exists, menu, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{Id: in.MenuId, AppId: in.AppId})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists || menu.Id != in.MenuId {
		return errc.ErrNotFound.MultiMsg("menu")
	}
	// check action exists
	actions, err := svc.d.FindActionConfig(ctx, &model.FindActionConfigReq{ActionId: in.ActionId, AppId: in.AppId})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	for _, v := range in.ActionId {
		hit := false
		for _, vv := range actions {
			if v == vv.Id {
				hit = true
				break
			}
		}
		if !hit {
			return errc.ErrNotFound.MultiMsg(fmt.Sprintf("action id %d", v))
		}
	}

	menu.Actions = new(types.IntSplitStr).Marshal(in.ActionId)

	if err := svc.d.UpdateMenuConfig(ctx, menu, []string{"actions"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

// GetMenuTreeDirs 获取菜单树目录结构
func (svc *Service) GetMenuTreeDirs(ctx context.Context, in *model.GetMenuTreeDirsReq, out *model.GetMenuTreeDirsResp) error {
	dirs, err := svc.menuTreeDirs(ctx, in)
	if err != nil {
		return err
	}
	out.Dirs = dirs
	return nil
}

func (svc *Service) GetMenuActions(ctx context.Context, in *model.GetMenuActionsReq, out *model.GetMenuActionsResp) error {
	exists, menu, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{
		Id:    in.MenuId,
		AppId: in.AppId,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("menuId")
	}

	actionId := menu.Actions.Unmarshal()
	if len(actionId) <= 0 {
		out.Actions = make(model.Actions, 0)
		return nil
	}
	actions, err := svc.d.FindActionConfig(ctx, &model.FindActionConfigReq{
		AppId:    in.AppId,
		ActionId: actionId,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	list := make(model.Actions, 0)
	for _, v := range actions {
		list = append(list, &model.Action{
			Id:     v.Id,
			AppId:  v.AppId,
			Name:   v.Name,
			Path:   v.Path,
			Method: v.Method,
		})
	}
	out.Actions = list
	return nil
}

func (svc *Service) menuTreeDirs(ctx context.Context, in *model.GetMenuTreeDirsReq) (model.MenuTreeDirs, error) {
	menus, err := svc.d.FindMenuConfig(ctx, &model.FindMenuConfigReq{
		AppId:    in.AppId,
		ParentId: -1, // all menus
	})
	if err != nil {
		return nil, errc.ErrInternalErr.MultiErr(err)
	}
	dirsMap := make(map[int64]model.MenuTreeDirs, 0)
	for _, v := range menus {
		dir := &model.MenuTreeDir{
			Id:       v.Id,
			AppId:    v.AppId,
			Name:     v.Name,
			Memo:     v.Memo,
			ParentId: v.ParentId,
			Status:   v.Status,
			Sequence: v.Sequence,
			Children: make(model.MenuTreeDirs, 0),
		}
		val, ok := dirsMap[v.ParentId]
		if !ok {
			dirsMap[v.ParentId] = []*model.MenuTreeDir{dir}
			continue
		}
		dirsMap[v.ParentId] = append(val, dir)
	}
	rootDirs := make(model.MenuTreeDirs, 0)
	if rootDir, ok := dirsMap[0]; ok {
		rootDirs = append(rootDirs, rootDir...)
	}
	sort.Sort(rootDirs)

	for _, root := range rootDirs {
		var findChildren func(root *model.MenuTreeDir, parentId int64)
		findChildren = func(root *model.MenuTreeDir, parentId int64) {
			children, ok := dirsMap[parentId]
			if !ok {
				return
			}
			sort.Sort(children)
			root.Children = children
			for _, v := range children {
				findChildren(v, v.Id)
			}
		}
		findChildren(root, root.Id)
	}

	//tests.PrintBeautifyJSON(rootDirs)

	return rootDirs, nil
}
