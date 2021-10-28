package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [APP配置列表]
// @Description
// @Tags RBAC APP配置
// @Security ApiKeyAuth
// @Produce json
// @Param Request query model.ListAppConfigReq true "request param"
// @Success 200 {object} model.ListAppConfig "success"
// @Router /rbac/v1/app/list [get]
func listAppConfig(c *gin.Context) {
	in := &model.ListAppConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.ListAppConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [APP配置筛选列表]
// @Description 用于Select组件
// @Tags RBAC APP配置
// @Accept json
// @Produce json
// @Param Request body model.SelectAppConfigReq true "request param"
// @Success 200 {object} model.SelectAppConfig "success"
// @Router /rbac/v1/app/select [get]
func selectAppConfig(c *gin.Context) {
	in := &model.SelectAppConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &typ.ListResp{}
	err := svc.SelectAppConfig(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}

// @Summary [APP配置创建]
// @Description
// @Tags RBAC APP配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.CreateAppConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/app/create [post]
func createAppConfig(c *gin.Context) {
	in := &model.CreateAppConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.CreateAppConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [APP配置更新]
// @Description
// @Tags RBAC APP配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.UpdateAppConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/app/update [post]
func updateAppConfig(c *gin.Context) {
	in := &model.UpdateAppConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.UpdateAppConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}

// @Summary [APP配置删除]
// @Description
// @Tags RBAC APP配置
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request body model.DelAppConfigReq true "request param"
// @Success 200 {object} ginutil.BaseResp "success"
// @Router /rbac/v1/app/delete [post]
func delAppConfig(c *gin.Context) {
	in := &model.DelAppConfigReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	err := svc.DelAppConfig(c.Request.Context(), in)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespSuccess(c)
}
