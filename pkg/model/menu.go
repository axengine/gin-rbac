package model

import (
	"github.com/bbdshow/admin-rabc/pkg/types"
	"time"
)

type MenuConfig struct {
	Id        int64             `xorm:"not null pk autoincr BIGINT(20)"`
	AppId     int64             `xorm:"not null BIGINT(20) comment('APP分组')"`
	Name      string            `xorm:"not null VARCHAR(128) comment('名称')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	ParentId  int64             `xorm:"not null BIGINT(20) comment('父ID')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	Actions   types.IntSliceStr `xorm:"not null TEXT comment('功能ID')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}

type ActionConfig struct {
	Id     int64  `xorm:"not null pk autoincr BIGINT(20)"`
	AppId  int64  `xorm:"not null BIGINT(20) comment('APP分组')"`
	Name   string `xorm:"not null VARCHAR(128) comment('名称')"`
	Path   string `xorm:"not null VARCHAR(256) comment('访问路径')"`
	Method int32  `xorm:"not null INT comment('GET=1 POST=2 PUT=3 DELETE=4')"`
}
