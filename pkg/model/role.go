package model

import (
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

type RoleConfig struct {
	Id        int64             `xorm:"not null pk autoincr INT(20)"`
	AppId     string            `xorm:"not null VARCHAR(6) comment('APP分组')"`
	Name      string            `xorm:"not null VARCHAR(128) comment('名称')"`
	IsRoot    int32             `xorm:"not null TINYINT(2) comment('ROOT 1-ROOT')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}

type RoleMenuAction struct {
	Id       int64 `xorm:"not null pk autoincr INT(20)"`
	RoleId   int64 `xorm:"not null BIGINT(20) comment('角色ID')"`
	MenuId   int64 `xorm:"not null BIGINT(20) comment('主菜单ID')"`
	ActionId int64 `xorm:"not null BIGINT(20) comment('功能ID')"`
}
