package service

import (
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/gin-rabc/pkg/conf"
	"github.com/bbdshow/gin-rabc/pkg/dao"
	"github.com/casbin/casbin/v2"
	"time"
)

type Service struct {
	cfg *conf.Config
	d   *dao.Dao

	enforce      *casbin.SyncedEnforcer
	enforceClose func()
}

func New(cfg *conf.Config) *Service {
	svc := &Service{
		cfg: cfg,
		d:   dao.New(cfg),
	}

	se, err := svc.initSyncEnforce()
	if err != nil {
		panic(fmt.Sprintf("casbin enforce %v", err))
	}
	svc.enforce = se
	return svc
}

func (svc *Service) Close() {
	if svc.d != nil {
		svc.d.Close()
	}
	if svc.enforceClose != nil {
		svc.enforceClose()
	}
}

func (svc *Service) initSyncEnforce() (*casbin.SyncedEnforcer, error) {
	se, err := casbin.NewSyncedEnforcer(casbinModel, NewCasbinAdapter(svc.d), svc.cfg.Casbin.Debug)
	if err != nil {
		return nil, errc.WithStack(err)
	}
	se.EnableEnforce(svc.cfg.Casbin.Enable)
	if svc.cfg.Casbin.Enable {
		se.StartAutoLoadPolicy(30 * time.Second)
		svc.enforceClose = func() {
			se.StopAutoLoadPolicy()
		}
	}
	return se, nil
}

func (svc *Service) GetSyncedEnforcer() *casbin.SyncedEnforcer {
	return svc.enforce
}
