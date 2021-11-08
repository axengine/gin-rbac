package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/bkit/util/icopier"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"sort"
	"strings"
)

func (svc *Service) ListActionConfig(ctx context.Context, in *model.ListActionConfigReq, out *typ.ListResp) error {
	c, records, err := svc.d.ListActionConfig(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	list := make(model.ListActionConfigs, 0, len(records))
	for _, v := range records {
		d := &model.ListActionConfig{}
		if err := icopier.Copy(d, v); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		d.UpdatedAt = v.UpdatedAt.Unix()
		list = append(list, d)
	}
	//sort.Sort(list)
	out.Count = c
	out.List = list

	return nil
}

func (svc *Service) FindActionConfig(ctx context.Context, in *model.FindActionConfigReq, out *model.FindActionConfigResp) error {
	if len(in.ActionId) <= 0 {
		out.Actions = make([]*model.ActionBase, 0)
		return nil
	}
	records, err := svc.d.FindActionConfig(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	list := make([]*model.ActionBase, 0, len(records))
	for _, v := range records {
		d := &model.ActionBase{}
		if err := icopier.Copy(d, v); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		list = append(list, d)
	}
	out.Actions = list
	return nil
}

func (svc *Service) CreateMenuConfig(ctx context.Context, in *model.CreateMenuConfigReq) error {
	r := &model.MenuConfig{
		AppId:    in.AppId,
		Name:     in.Name,
		Memo:     in.Memo,
		ParentId: in.ParentId,
		Status:   types.LimitNormal,
		Sequence: in.Sequence,
		Path:     in.Path,
		Typ:      in.Typ,
	}

	if in.ParentId != 0 {
		exists, _, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{Id: in.ParentId})
		if err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		if !exists {
			return errc.ErrParamInvalid.MultiMsg("parentId not found")
		}
	}

	if err := svc.d.CreateMenuConfig(ctx, r); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateMenuConfig(ctx context.Context, in *model.UpdateMenuConfigReq) error {
	d := &model.MenuConfig{
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
	if len(in.Path) > 0 {
		cols = append(cols, "path")
		d.Path = in.Path
	}
	if in.Typ > 0 {
		cols = append(cols, "typ")
		d.Typ = in.Typ
	}
	if in.ParentId > -1 {
		if in.ParentId != 0 {
			exists, _, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{Id: in.ParentId})
			if err != nil {
				return errc.ErrInternalErr.MultiErr(err)
			}
			if !exists {
				return errc.ErrParamInvalid.MultiMsg("parentId not found")
			}
		}
		cols = append(cols, "parent_id")
		d.ParentId = in.ParentId
	}
	if in.Sequence > -1 {
		cols = append(cols, "sequence")
		d.Sequence = in.Sequence
	}

	if err := svc.d.UpdateMenuConfig(ctx, d, cols); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) ImportActionConfig(ctx context.Context, in *model.ImportActionConfigReq) error {
	in.Method = strings.ToUpper(in.Method)
	if err := svc.d.ImportActionConfig(ctx, in); err != nil {
		return errc.WithStack(err)
	}
	return nil
}

func (svc *Service) CreateActionConfig(ctx context.Context, in *model.CreateActionConfigReq) error {
	in.Method = strings.ToUpper(in.Method)
	exists, _, err := svc.d.GetActionConfig(ctx, &model.GetActionConfigReq{
		AppId:  in.AppId,
		Path:   in.Path,
		Method: in.Method,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if exists {
		return errc.ErrParamInvalid.MultiMsg("action config exists")
	}

	r := &model.ActionConfig{
		AppId:  in.AppId,
		Name:   in.Name,
		Path:   in.Path,
		Method: in.Method,
		Status: types.LimitNormal,
	}

	if err := svc.d.CreateActionConfig(ctx, r); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	return nil
}

func (svc *Service) UpdateActionConfig(ctx context.Context, in *model.UpdateActionConfigReq) error {
	r := &model.ActionConfig{
		Id: in.Id,
	}
	cols := make([]string, 0)
	if len(in.Name) > 0 {
		cols = append(cols, "name")
		r.Name = in.Name
	}
	if in.Path != "" {
		cols = append(cols, "path")
		r.Path = in.Path
	}

	if in.Method != "" {
		cols = append(cols, "method")
		r.Method = in.Method
	}

	if in.Status > 0 {
		cols = append(cols, "status")
		r.Status = in.Status
	}

	if err := svc.d.UpdateActionConfig(ctx, r, cols); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) DelActionConfig(ctx context.Context, in *model.DelActionConfigReq) error {
	menus, err := svc.d.FindMenuConfig(ctx, &model.FindMenuConfigReq{ActionId: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	for _, v := range menus {
		hit := false
		for _, vv := range v.Actions.Unmarshal() {
			if vv == in.Id {
				hit = true
				break
			}
		}
		if hit {
			return errc.ErrParamInvalid.MultiMsg("action is bound to menu , not allow delete, should setting status")
		}
	}

	if err := svc.d.DelActionConfig(ctx, in.Id); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateMenuConfigAction(ctx context.Context, in *model.UpdateMenuConfigActionReq) error {
	exists, menu, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{Id: in.MenuId})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists || menu.Id != in.MenuId {
		return errc.ErrNotFound.MultiMsg("menu")
	}
	// remove repeat
	actionsIds := make([]int64, 0, len(in.ActionId))
	for _, v := range in.ActionId {
		hit := false
		for _, vv := range actionsIds {
			if v == vv {
				hit = true
				break
			}
		}
		if !hit {
			actionsIds = append(actionsIds, v)
		}
	}
	in.ActionId = actionsIds

	// check action exists
	actions, err := svc.d.FindActionConfig(ctx, &model.FindActionConfigReq{ActionId: in.ActionId, AppId: menu.AppId})
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

func (svc *Service) DelMenuConfig(ctx context.Context, in *model.DelMenuConfigReq) error {
	// if have children, not allow del
	children, err := svc.d.FindMenuConfig(ctx, &model.FindMenuConfigReq{
		ParentId: in.Id,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if len(children) > 0 {
		return errc.ErrParamInvalid.MultiMsg("have children menu, not allow delete")
	}

	if err := svc.d.DelMenuConfig(ctx, in.Id); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

// GetMenuTreeDirs 获取菜单树目录结构
func (svc *Service) GetMenuTreeDirs(ctx context.Context, in *model.GetMenuTreeDirsReq, out *model.GetMenuTreeDirsResp) error {
	menus, err := svc.d.FindMenuConfig(ctx, &model.FindMenuConfigReq{
		AppId:    in.AppId,
		ParentId: -1, // all menus
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	dirs, err := svc.menuTreeDirs(menus)
	if err != nil {
		return err
	}
	out.Dirs = dirs
	return nil
}

func (svc *Service) GetMenuActions(ctx context.Context, in *model.GetMenuActionsReq, out *model.GetMenuActionsResp) error {
	exists, menu, err := svc.d.GetMenuConfig(ctx, &model.GetMenuConfigReq{
		Id: in.MenuId,
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
		AppId:    menu.AppId,
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

func (svc *Service) menuTreeDirs(menus []*model.MenuConfig) (model.MenuTreeDirs, error) {
	dirsMap := make(map[int64]model.MenuTreeDirs, 0)
	for _, v := range menus {
		dir := &model.MenuTreeDir{
			Id:       v.Id,
			AppId:    v.AppId,
			Name:     v.Name,
			Typ:      v.Typ,
			Memo:     v.Memo,
			ParentId: v.ParentId,
			Status:   v.Status,
			Sequence: v.Sequence,
			Path:     v.Path,
			Actions:  v.Actions.Unmarshal(),
			Children: make(model.MenuTreeDirs, 0),
		}
		val, ok := dirsMap[v.ParentId]
		if !ok {
			dirsMap[v.ParentId] = model.MenuTreeDirs{dir}
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
