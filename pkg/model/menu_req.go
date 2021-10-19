package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/typ"
)

type ListMenuConfigReq struct {
}

type FindMenuConfigReq struct {
	AppId    int64
	ParentId int64
}
type GetMenuConfigReq struct {
	Id    int64
	AppId int64
}

type CreateMenuConfigReq struct {
	AppId    int64   `json:"appId"`
	Name     string  `json:"name"`
	Memo     string  `json:"memo"`
	ParentId int64   `json:"parentId"`
	Sequence int     `json:"sequence"`
	Actions  []int64 `json:"actions"`
}

type UpsertActionConfigReq struct {
	AppId  int64  `json:"appId" binding:"required,min=1"`
	Name   string `json:"name" binding:"required,gte=1,gte=128"`
	Path   string `json:"path" binding:"required,gte=1,gte=256"`
	Method string `json:"method" binding:"required,upper"`
}

type UpdateMenuActionReq struct {
	AppId    int64   `json:"appId"`
	MenuId   int64   `json:"menuId"`
	ActionId []int64 `json:"actionId"`
}

type ListActionConfigReq struct {
	typ.PageReq
}

type FindActionConfigReq struct {
	AppId    int64
	ActionId []int64
}

type GetMenuTreeDirsReq struct {
	AppId int64
}

type GetMenuTreeDirsResp struct {
	Dirs MenuTreeDirs `json:"dirs"`
}

type GetMenuActionsReq struct {
	AppId  int64 `json:"appId" form:"appId"`
	MenuId int64 `json:"menuId" form:"menuId"`
}

type GetMenuActionsResp struct {
	Actions Actions `json:"actions"`
}

type MenuTrees []*MenuTree

type MenuTree struct {
	AppId    int64             `json:"appId"`
	Name     string            `json:"name"`
	Memo     string            `json:"memo"`
	ParentId int64             `json:"parentId"`
	Status   types.LimitStatus `json:"status"`
	Sequence int               `json:"sequence"`
	Actions  Actions           `json:"actions"`
	Children MenuTrees         `json:"children"`
}

type Actions []*Action
type Action struct {
	Id     int64  `json:"id"`
	AppId  int64  `json:"appId"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type MenuTreeDirs []*MenuTreeDir

func (dirs MenuTreeDirs) Len() int           { return len(dirs) }
func (dirs MenuTreeDirs) Swap(i, j int)      { dirs[i], dirs[j] = dirs[j], dirs[i] }
func (dirs MenuTreeDirs) Less(i, j int) bool { return dirs[i].Sequence < dirs[j].Sequence }

type MenuTreeDir struct {
	Id       int64             `json:"id"`
	AppId    int64             `json:"appId"`
	Name     string            `json:"name"`
	Memo     string            `json:"memo"`
	ParentId int64             `json:"parentId"`
	Status   types.LimitStatus `json:"status"`
	Sequence int               `json:"sequence"`
	Children MenuTreeDirs      `json:"children"`
}
