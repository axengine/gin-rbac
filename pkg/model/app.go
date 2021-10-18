package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"time"
)

type AppConfig struct {
	Id        int64             `xorm:"not null pk autoincr INT(20)"`
	Name      string            `xorm:"not null VARCHAR(128) comment('名称')"`
	AccessKey string            `xorm:"not null VARCHAR(16) comment('AccessKey')"`
	SecretKey string            `xorm:"not null VARCHAR(32) comment('SecretKey')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}
