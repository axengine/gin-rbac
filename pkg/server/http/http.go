package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/runner"
	"github.com/bbdshow/gin-rabc/pkg/conf"
	"github.com/bbdshow/gin-rabc/pkg/service"
)

var (
	svc *service.Service
	cfg *conf.Config
)

func NewAdminHttpServer(c *conf.Config, s *service.Service) runner.Server {
	svc = s
	cfg = c

	midFlag := ginutil.MStd
	if cfg.Release() {
		midFlag = ginutil.MRelease | ginutil.MTraceId | ginutil.MRecoverLogger
	}
	httpHandler := ginutil.DefaultEngine(midFlag)
	RegisterAdminRouter(httpHandler)

	return runner.NewHttpServer(httpHandler)
}
