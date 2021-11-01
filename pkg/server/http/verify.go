package http

import (
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/gin-gonic/gin"
)

// @Summary [HTTP验证权限]
// @Description http请求验证是否拥有权限， 通过 accessToken 和要验证的 Path Method, 请求需签名
// @Tags RBAC HTTP验证权限
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Request query model.RBACEnforceReq true "request param"
// @Success 200 {object} model.RBACEnforceResp "success"
// @Router /rbac/v1/verify/enforce [post]
func rbacEnforce(c *gin.Context) {
	in := &model.RBACEnforceReq{}
	if err := ginutil.ShouldBind(c, in); err != nil {
		ginutil.RespErr(c, err)
		return
	}
	out := &model.RBACEnforceResp{}
	err := svc.RBACEnforce(c.Request.Context(), in, out)
	if err != nil {
		ginutil.RespErr(c, err)
		return
	}
	ginutil.RespData(c, out)
}
