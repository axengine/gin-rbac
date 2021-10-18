package service

import (
	"github.com/bbdshow/admin-rabc/pkg/conf"
	"github.com/bbdshow/admin-rabc/pkg/dao"
)

type Service struct {
	cfg *conf.Config
	d   *dao.Dao
}

func New(cfg *conf.Config) *Service {
	svc := &Service{
		cfg: cfg,
		d:   dao.New(cfg),
	}

	return svc
}

func (svc *Service) Close() {
	if svc.d != nil {
		svc.d.Close()
	}
}
