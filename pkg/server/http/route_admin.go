package http

import "github.com/gin-gonic/gin"

func RegisterAdminRouter(e *gin.Engine) {
	e.POST("/rbac/login", loginAccount)

	e.POST("/rbac/loginOut", MidAccessTokenVerify(), loginOutAccount)

	rbac := e.Group("/rbac/v1")

	app := rbac.Group("/app").Use(MidAccessTokenVerify())
	app.GET("/list", listAppConfig)
	app.POST("/create", createAppConfig)
	app.POST("/update", updateAppConfig)

	rbac.GET("/app/select", selectAppConfig)

	action := rbac.Group("/action").Use(MidAccessTokenVerify())
	action.GET("/list", listActionConfig)
	action.POST("/upsert", upsertActionConfig)

	account := rbac.Group("/account").Use(MidAccessTokenVerify())
	account.GET("/list", listAccount)
	account.POST("/create", createAccount)
	account.POST("/pwd/update", updateAccountPassword)
	account.POST("/role/update", updateAccountRole)

	role := rbac.Group("/role").Use(MidAccessTokenVerify())
	role.GET("/list", listRoleConfig)
	role.POST("/create", createRoleConfig)
	role.POST("/update", updateRoleConfig)

}
