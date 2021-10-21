package model

import "github.com/bbdshow/bkit/typ"

type GetRoleConfigReq struct {
	Id int64
}

type CreateRoleConfigReq struct {
	AppId  string `json:"appId" binding:"required,len=6"`
	Name   string `json:"name" binding:"required,gte=1,lte=128"`
	IsRoot int32  `json:"isRoot"`
	Memo   string `json:"memo"`
}

type UpdateRoleConfigReq struct {
	typ.IdReq
	CreateRoleConfigReq // appId 不支持更改
}

type UpsertRoleMenuActionReq struct {
	AppId string
}
