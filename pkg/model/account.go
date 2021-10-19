package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"time"
)

type Account struct {
	Id        int64             `xorm:"not null pk autoincr INT(20)"`
	AppId     int64             `xorm:"not null BIGINT(20) comment('APP分组')"`
	Nickname  string            `xorm:"not null VARCHAR(64) comment('昵称')"`
	Username  string            `xorm:"not null VARCHAR(64) comment('账号名')"`
	Password  string            `xorm:"not null VARCHAR(64) comment('密码')"`
	Salt      string            `xorm:"not null VARCHAR(6) comment('盐')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	Roles     types.IntSplitStr `xorm:"not null TEXT comment('角色ID')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}
