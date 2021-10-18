package http

import "github.com/gin-gonic/gin"

func RegisterAdminRouter(e *gin.Engine) {
	e.POST("/login")
	//
	//// 中间件都可以注册在这里
	//jwt := ginutil.JWTAuthVerify(cfg.Admin.AuthEnable)
	//
	admin := e.Group("/v1/admin")
	admin.GET("")
}
