package model

import (
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

type ListAccountReq struct {
	AppId    string            `json:"appId" form:"appId"`
	Nickname string            `json:"nickname" form:"nickname"`
	Username string            `json:"username" form:"username"`
	Status   types.LimitStatus `json:"status" form:"status"`
	typ.PageReq
}

type ListAccount struct {
	Id           int64             `json:"id"`
	AppName      string            `json:"appName"`
	AppId        string            `json:"appId"`
	Nickname     string            `json:"nickname"`
	Username     string            `json:"username"`
	PwdWrong     int               `json:"pwdWrong"`
	LoginLock    int64             `json:"loginLock"`
	TokenExpired int64             `json:"tokenExpired"`
	Memo         string            `json:"memo"`
	Status       types.LimitStatus `json:"status"`
	Roles        []RoleBase        `json:"roles"`
	UpdatedAt    int64             `json:"updatedAt"`
	CreatedAt    int64             `json:"createdAt"`
}

type RoleBase struct {
	Id     int64             `json:"id"`
	Name   string            `json:"name"`
	Status types.LimitStatus `json:"status"`
}

type GetAccountReq struct {
	Id       int64
	AppId    string
	Username string
	Token    string
}

type FindAccountReq struct {
	AppId  string
	Status types.LimitStatus
}

type CreateAccountReq struct {
	AppId    string `json:"appId" binding:"required,len=6"`
	Nickname string `json:"nickname" binding:"required,lte=64"`
	Username string `json:"username" binding:"required,lte=64"`
	Password string `json:"pwd" binding:"required,lte=32"`
}

type DelAccountReq struct {
	typ.IdReq
}

type LoginAccountReq struct {
	AppId    string `json:"appId" binding:"required,len=6"`
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

type VerifyAccountTokenResp struct {
	Id        int64
	AppId     string
	Nickname  string
	Username  string
	LoginLock int64
	Status    types.LimitStatus
	IsRoot    int32
	Roles     types.IntSplitStr
	CreatedAt time.Time
}
