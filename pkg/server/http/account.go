package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [账户配置列表]
// @Description
// @Tags RBAC 账户配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.ListAccountReq true "request param"
// @Success 200 {object} model.ListAccount "success"
// @Router /rbac/v1/account/list [get]
func listAccount(c *gin.Context) {
	in := &model.ListAccountReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.ListAccount(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [账户配置创建]
// @Description
// @Tags RBAC 账户配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateAccountReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/account/create [post]
func createAccount(c *gin.Context) {
	in := &model.CreateAccountReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateAccount(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [账户配置删除]
// @Description
// @Tags RBAC 账户配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.DelAccountReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/account/delete [post]
func delAccount(c *gin.Context) {
	in := &model.DelAccountReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.DelAccount(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [账户配置密码更改]
// @Description
// @Tags RBAC 账户配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateAccountPasswordReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/account/pwd/update [post]
func updateAccountPassword(c *gin.Context) {
	in := &model.UpdateAccountPasswordReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateAccountPassword(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [账户配置角色更改]
// @Description
// @Tags RBAC 账户配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateAccountRoleReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/account/role/update [post]
func updateAccountRole(c *gin.Context) {
	in := &model.UpdateAccountRoleReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateAccountRole(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}
