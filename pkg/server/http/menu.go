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
