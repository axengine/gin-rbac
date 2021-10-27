package main

import (
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/bkit/runner"
	"github.com/bbdshow/gin-rabc/pkg/conf"
	"github.com/bbdshow/gin-rabc/pkg/server/http"
	"github.com/bbdshow/gin-rabc/pkg/service"
	"go.uber.org/zap"
	"log"

	_ "github.com/bbdshow/gin-rabc/docs"
)

// @title gin rbac
// @version 1.0.0
// @description gin rbac manage API

// @host API_HOST:49000
// @BasePath /
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
