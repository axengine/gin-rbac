package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/typ"
)

type ListAppConfigReq struct {
	Name   string `form:"name"`
	Status int    `form:"status"`
	typ.PageReq
}

type GetAppConfigReq struct {
	Id        int64
	AccessKey string
}

type GetAppConfigResp struct {
	Id        int64
	Name      string            `json:"name"`
	AccessKey string            `json:"accessKey"` // 访问KEY
	SecretKey string            `json:"secretKey"` // 加密KEY
	Status    types.LimitStatus `json:"status"`    // 状态 1-正常 2-限制
	Memo      string            `json:"memo"`      // 备注
}

type ListAppConfig struct {
	Id        int64             `json:"id"`
	Name      string            `json:"name"`      // APP名
	AccessKey string            `json:"accessKey"` // 访问KEY
	SecretKey string            `json:"secretKey"` // 加密KEY
	Status    types.LimitStatus `json:"status"`    // 状态 1-正常 2-限制
	Memo      string            `json:"memo"`      // 备注
	UpdatedAt int64             `json:"updatedAt"`
	CreatedAt int64             `json:"createdAt"`
}

type CreateAppConfigReq struct {
	Name string `json:"name" binding:"required,gte1,lte=128"`
	Memo string `json:"memo" binding:"required,lte=128"`
}

type UpdateAppConfigReq struct {
	typ.IdReq
	SecretKey bool `json:"secretKey"` // true = 重置加密KEY
	CreateAppConfigReq
	Status types.LimitStatus `json:"status"` //状态 1-正常 2-限制
}
