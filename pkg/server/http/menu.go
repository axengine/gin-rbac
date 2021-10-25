package http

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/gin-gonic/gin"
)

// @Summary [功能配置列表]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.ListActionConfigReq true "request param"
// @Success 200 {object} model.ListActionConfig "success"
// @Router /rbac/v1/action/list [get]
func listActionConfig(c *gin.Context) {
	in := &model.ListActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.ListActionConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [功能配置创建/更新]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpsertActionConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/action/upsert [post]
func upsertActionConfig(c *gin.Context) {
	in := &model.UpsertActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpsertActionConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [菜单配置树]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.GetMenuTreeDirsReq true "request param"
// @Success 200 {object} model.GetMenuTreeDirsResp "success"
// @Router /rbac/v1/menu/tree [get]
func treeMenuConfig(c *gin.Context) {
	in := &model.GetMenuTreeDirsReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.GetMenuTreeDirsResp{}
	err := svc.GetMenuTreeDirs(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [菜单配置功能]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.GetMenuActionsReq true "request param"
// @Success 200 {object} model.GetMenuActionsResp "success"
// @Router /rbac/v1/menu/actions [get]
func menuActions(c *gin.Context) {
	in := &model.GetMenuActionsReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.GetMenuActionsResp{}
	err := svc.GetMenuActions(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [菜单配置创建]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateMenuConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/menu/create [post]
func createMenuConfig(c *gin.Context) {
	in := &model.CreateMenuConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateMenuConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [菜单配置更新]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateMenuConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/menu/update [post]
func updateMenuConfig(c *gin.Context) {
	in := &model.UpdateMenuConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateMenuConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [菜单配置功能更新]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateMenuConfigActionReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/menu/action/update [post]
func updateMenuConfigAction(c *gin.Context) {
	in := &model.UpdateMenuConfigActionReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateMenuConfigAction(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}
