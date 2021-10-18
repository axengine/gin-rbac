package service

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/conf"
	"os"
	"testing"
)

var (
	svc *Service
	ctx = context.Background()
)

func TestMain(m *testing.M) {
	if err := conf.InitConf("../../configs/config.toml"); err != nil {
		panic(err)
	}
	svc = New(conf.Conf)
	os.Exit(m.Run())
}
