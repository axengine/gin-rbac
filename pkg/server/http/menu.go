package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [功能配置列表]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Produce json
// @Param Request query model.ListActionConfigReq true "request param"
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

// @Summary [功能配置查询]
// @Description 根据功能ID，查询基础信息
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.FindActionConfigReq true "request param"
// @Success 200 {object} model.FindActionConfigResp "success"
// @Router /rbac/v1/action/find [post]
func findActionConfig(c *gin.Context) {
	in := &model.FindActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.FindActionConfigResp{}
	err := svc.FindActionConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [功能配置创建]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateActionConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/action/create [post]
func createActionConfig(c *gin.Context) {
	in := &model.CreateActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateActionConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [功能配置更新]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateActionConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/action/update [post]
func updateActionConfig(c *gin.Context) {
	in := &model.UpdateActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateActionConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [功能配置删除]
// @Description
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.DelActionConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/action/delete [post]
func delActionConfig(c *gin.Context) {
	in := &model.DelActionConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.DelActionConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [功能配置导入Swagger]
// @Description 导入Swagger JSON 文件
// @Tags RBAC 功能配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.SwaggerJSONToActionsReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/action/import [post]
func importSwaggerToActions(c *gin.Context) {
	in := &model.SwaggerJSONToActionsReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.SwaggerJSONToActions(c.Request.Context(), in)
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
// @Produce json
// @Param Request query model.GetMenuTreeDirsReq true "request param"
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
// @Produce json
// @Param Request query model.GetMenuActionsReq true "request param"
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

// @Summary [菜单配置删除]
// @Description
// @Tags RBAC 菜单配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.DelMenuConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/menu/delete [post]
func delMenuConfig(c *gin.Context) {
	in := &model.DelMenuConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.DelMenuConfig(c.Request.Context(), in)
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
