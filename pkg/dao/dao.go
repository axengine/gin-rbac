package dao

import (
	"github.com/bbdshow/admin-rabc/pkg/conf"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/caches"
	"github.com/bbdshow/bkit/db/mysql"
)

type Dao struct {
	cfg   *conf.Config
	mysql *mysql.Xorm

	memCache caches.Cacher
}

func New(cfg *conf.Config) *Dao {
	d := &Dao{
		cfg:   cfg,
		mysql: mysql.NewXorm(cfg.Mysql),

		memCache: caches.NewLRUMemory(10000),
	}
	d.Sync2()

	return d
}

func (d *Dao) Close() {
	if d.mysql != nil {
		_ = d.mysql.Close()
	}
}

func (d *Dao) Sync2() {
	if !d.cfg.Release() {
		err := d.mysql.Sync2(
			new(model.AppConfig),
			new(model.MenuConfig),
			new(model.ActionConfig),
			new(model.RoleConfig),
			new(model.RoleMenuAction),
			new(model.Account),
		)
		if err != nil {
			panic(err)
		}
	}
}
