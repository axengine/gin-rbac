package http

import "github.com/gin-gonic/gin"

func RegisterAdminRouter(e *gin.Engine) {
	e.POST("/rbac/login", loginAccount)
	e.POST("/rbac/loginOut", loginOutAccount)

	rbac := e.Group("/rbac/v1")

	app := rbac.Group("/app")
	app.GET("/list", listAppConfig)
	app.POST("/create", createAppConfig)
	app.POST("/update", updateAppConfig)

	action := rbac.Group("/action")
	action.GET("/list", listActionConfig)
	action.POST("/upsert", upsertActionConfig)

	account := rbac.Group("/account")
	account.GET("/list", listAccount)
	account.POST("/create", createAccount)
	account.POST("/pwd/update", updateAccountPassword)
	account.POST("/role/update", updateAccountRole)

	role := rbac.Group("/role")
	role.GET("/list", listRoleConfig)
	role.POST("/create", createRoleConfig)
	role.POST("/update", updateRoleConfig)

}
