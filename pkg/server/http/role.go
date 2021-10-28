package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [角色配置列表]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Produce json
// @Param Request query model.ListRoleConfigReq true "request param"
// @Success 200 {object} model.ListRoleConfig "success"
// @Router /rbac/v1/role/list [get]
func listRoleConfig(c *gin.Context) {
	in := &model.ListRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.ListRoleConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [获取角色菜单功能]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Produce json
// @Param Request query model.GetRoleMenuActionReq true "request param"
// @Success 200 {object} model.GetRoleMenuActionResp "success"
// @Router /rbac/v1/role/action [get]
func getRoleMenuAction(c *gin.Context) {
	in := &model.GetRoleMenuActionReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.GetRoleMenuActionResp{}
	err := svc.GetRoleMenuAction(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [角色配置创建]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateRoleConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/create [post]
func createRoleConfig(c *gin.Context) {
	in := &model.CreateRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateRoleConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [角色配置更新]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateRoleConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/update [post]
func updateRoleConfig(c *gin.Context) {
	in := &model.UpdateRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateRoleConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [角色配置删除]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateRoleConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/delete [post]
func delRoleConfig(c *gin.Context) {
	in := &model.DelRoleConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.DelRoleConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [角色配置功能绑定]
// @Description
// @Tags RBAC 角色配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpsertRoleMenuActionReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/role/action/upsert [post]
func upsertRoleMenuAction(c *gin.Context) {
	in := &model.UpsertRoleMenuActionReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpsertRoleMenuAction(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}
