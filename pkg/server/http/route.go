package http

import (
	"github.com/bbdshow/gin-rabc/pkg/server/http/middleware"
	"github.com/gin-gonic/gin"
)

func midAccessTokenVerify() gin.HandlerFunc {
	return middleware.AccessTokenVerify(cfg.Admin.AuthEnable, svc.VerifyAccountToken)
}

func MidRBACEnforce() gin.HandlerFunc {
	enable := cfg.Admin.AuthEnable && cfg.Casbin.Enable
	return middleware.RBACEnforce(enable, svc.GetSyncedEnforcer())
}

func RegisterAdminRouter(e *gin.Engine) {
	e.POST("/rbac/login", loginAccount)

	e.POST("/rbac/loginOut", midAccessTokenVerify(), loginOutAccount)

	rbac := e.Group("/rbac/v1")

	app := rbac.Group("/app").Use(midAccessTokenVerify()).Use(MidRBACEnforce())
	app.GET("/list", listAppConfig)
	app.POST("/create", createAppConfig)
	app.POST("/update", updateAppConfig)
	app.POST("/delete", delAppConfig)

	rbac.GET("/app/select", selectAppConfig)

	action := rbac.Group("/action").Use(midAccessTokenVerify()).Use(MidRBACEnforce())
	action.GET("/list", listActionConfig)
	action.POST("/upsert", upsertActionConfig)
	action.POST("/delete", delActionConfig)
	action.POST("/find", findActionConfig)
	action.POST("/import", importSwaggerToActions)

	menu := rbac.Group("/menu").Use(midAccessTokenVerify()).Use(MidRBACEnforce())
	menu.GET("/tree", treeMenuConfig)
	menu.GET("/actions", menuActions)
	menu.POST("/create", createMenuConfig)
	menu.POST("/update", updateMenuConfig)
	menu.POST("/delete", delMenuConfig)
	menu.POST("/action/update", updateMenuConfigAction)

	account := rbac.Group("/account").Use(midAccessTokenVerify()).Use(MidRBACEnforce())
	account.GET("/list", listAccount)
	account.POST("/create", createAccount)
	account.POST("/delete", delAccount)
	account.POST("/pwd/update", updateAccountPassword)
	account.POST("/role/update", updateAccountRole)
	rbac.GET("/account/menu/auth", midAccessTokenVerify(), getAccountMenuAuth)

	role := rbac.Group("/role").Use(midAccessTokenVerify()).Use(MidRBACEnforce())
	role.GET("/list", listRoleConfig)
	role.POST("/create", createRoleConfig)
	role.POST("/update", updateRoleConfig)
	role.POST("/delete", delRoleConfig)
	role.GET("/action", getRoleMenuAction)
	role.POST("/action/upsert", upsertRoleMenuAction)
}
