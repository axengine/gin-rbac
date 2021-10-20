package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/typ"
)

type ListAccountReq struct {
	AppId    int64             `json:"appId" form:"appId"`
	Nickname string            `json:"nickname" form:"nickname"`
	Username string            `json:"username" form:"username"`
	Status   types.LimitStatus `json:"status" form:"status"`
	typ.PageReq
}

type ListAccount struct {
	AppName      string            `json:"appName"`
	AppId        int64             `json:"appId"`
	Nickname     string            `json:"nickname"`
	PwdWrong     int               `json:"pwdWrong"`
	LoginLock    int64             `json:"loginLock"`
	TokenExpired int64             `json:"tokenExpired"`
	Memo         string            `json:"memo"`
	Status       types.LimitStatus `json:"status"`
	Roles        []string          `json:"roles"`
	CreatedAt    int64             `json:"createdAt"`
}

type GetAccountReq struct {
	Id       int64
	AppId    int64
	Username string
}

type CreateAccountReq struct {
	AppId    int64  `json:"appId" binding:"required,min=1"`
	Nickname string `json:"nickname" binding:"required,lte=64"`
	Username string `json:"username" binding:"required,lte=64"`
	Password string `json:"pwd" binding:"required,lte=32"`
}

type LoginAccountReq struct {
	AppId    int64  `json:"appId" binding:"required,min=1"`
	Username string `json:"username" binding:"required,lte=64"`
	Password string `json:"password" binding:"required,len=32"`
}

type LoginAccountResp struct {
	Token        string `json:"token"`
	TokenExpired int64  `json:"tokenExpired"`
}

type LoginOutAccountReq struct {
	Token string `json:"-"`
}

type UpdateAccountPasswordReq struct {
	typ.IdReq
	Password string `json:"pwd" binding:"required,lte=32"`
}

type UpdateAccountRoleReq struct {
	typ.IdReq
	Roles []int64 `json:"roles" binding:"required"`
}
