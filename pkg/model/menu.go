package model

import (
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

type MenuConfig struct {
	Id        int64             `xorm:"not null pk autoincr BIGINT(20)"`
	AppId     string            `xorm:"not null VARCHAR(6) comment('APP分组')"`
	Name      string            `xorm:"not null VARCHAR(128) comment('名称')"`
	Typ       int               `xorm:"not null TINYINT(2) comment('分类 1-菜单 2-分组')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	ParentId  int64             `xorm:"not null BIGINT(20) comment('父ID')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	Sequence  int               `xorm:"not null default 0 INT comment('序号')"`
	Path      string            `xorm:"not null VARCHAR(255) comment('路径')"`
	Actions   types.IntSplitStr `xorm:"not null TEXT comment('功能ID')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}

type ActionConfig struct {
	Id        int64             `xorm:"not null pk autoincr BIGINT(20)"`
	AppId     string            `xorm:"not null VARCHAR(6) unique(appid_path_method) comment('APP分组')"`
	Name      string            `xorm:"not null VARCHAR(128) index comment('名称')"`
	Path      string            `xorm:"not null VARCHAR(255) unique(appid_path_method) comment('访问路径')"`
	Method    string            `xorm:"not null VARCHAR(10) unique(appid_path_method) comment('GET POST PUT DELETE')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	UpdatedAt time.Time         `xorm:"updated"`
}
