package main

import (
	"github.com/bbdshow/admin-rabc/pkg/conf"
	"github.com/bbdshow/admin-rabc/pkg/server/http"
	"github.com/bbdshow/admin-rabc/pkg/service"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/bkit/runner"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := conf.InitConf(); err != nil {
		panic(err)
	}
	logs.InitQezap(&conf.Conf.Logger)
	defer logs.Qezap.Close()

	logs.Qezap.Info("init", zap.Any("config", conf.Conf.EraseSensitive()))

	svc := service.New(conf.Conf)
	defer svc.Close()

	if err := runner.Run(http.NewAdminHttpServer(conf.Conf, svc), func() error {
		// dealloc
		return nil
	}, runner.WithListenAddr(conf.Conf.Admin.HttpListenAddr)); err != nil {
		log.Printf("runner exit: %v\n", err)
	}
}
