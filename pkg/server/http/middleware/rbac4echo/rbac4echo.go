package rbac4echo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/axengine/utils/httpx"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// RABC Middleware for https://github.com/axengine/gin-rbac
const (
	RbacAccessTokenHeader = "X-Rbac-Access-Token"
	RbacContextAccount    = "X-Rbac-Context-Account"
)

type RBACEnforceReq struct {
	// 在RBAC登录成功获得的TOKEN
	AccessToken string `json:"accessToken" binding:"required,len=32"`
	// 请求URI路径
	Path string `json:"path" binding:"required"`
	// 请求HTTP方法
	Method string `json:"method" binding:"required"`
}

type RBACEnforceResp struct {
	AccountId  int64  `json:"accountId"`
	AppId      string `json:"appId"`
	Nickname   string `json:"nickname"`
	IsRoot     int32  `json:"isRoot"`
	ActionPass bool   `json:"actionPass"` // false-无权限
}

func (r *RBACEnforceResp) Marshal() string {
	byt, _ := json.Marshal(r)
	return string(byt)
}

func (r *RBACEnforceResp) Unmarshal(v string) (*RBACEnforceResp, error) {
	err := json.Unmarshal([]byte(v), r)
	return r, err
}

func RBACEnforce(enable bool, domain, accessKey, secretKey string, skipPrefixPaths ...string) echo.MiddlewareFunc {
	signer := httpx.NewAPISign(accessKey, secretKey, 60, httpx.HmacSha256Hex)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !enable {
				return next(c)
			}
			path := c.Path()
			method := strings.ToUpper(c.Request().Method)
			// skip
			for _, p := range skipPrefixPaths {
				if strings.HasPrefix(path, p) {
					return next(c)
				}
			}
			accessToken := c.Request().Header.Get(RbacAccessTokenHeader)
			if len(accessToken) != 32 {
				return c.JSON(http.StatusForbidden, "Require Authorization")
			}

			in := &RBACEnforceReq{
				AccessToken: accessToken,
				Path:        path,
				Method:      method,
			}
			out := &struct {
				Code    int             `json:"code"`
				Message string          `json:"message"`
				Data    RBACEnforceResp `json:"data"`
			}{}
			body, _ := json.Marshal(in)

			req, err := http.NewRequest("POST",
				fmt.Sprintf("%s/rbac/v1/verify/enforce", domain),
				bytes.NewReader(body))
			if err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			}
			req.Header.Set("Content-Type", "application/json")

			// 接口签名
			deadline := time.Now().Add(time.Minute).Unix()
			if signature, err := signer.Sign(req, deadline); err != nil {
				return c.JSON(http.StatusForbidden, err.Error())
			} else {
				req.Header.Set("Authorization", fmt.Sprintf("%s:%s:%d", accessKey, signature, deadline))
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			defer resp.Body.Close()
			if bz, err := io.ReadAll(resp.Body); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			} else if err := json.Unmarshal(bz, &out); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			if out.Code != 0 {
				return c.JSON(http.StatusForbidden, out.Message)
			}
			c.Set(RbacContextAccount, out.Data.Marshal())

			if !out.Data.ActionPass {
				return c.JSON(http.StatusForbidden, out.Message)
			}
			return next(c)
		}
	}
}

func GetCtxRbacContextAccount(ctx echo.Context) (*RBACEnforceResp, bool) {
	v := ctx.Request().Header.Get(RbacContextAccount)
	if len(v) == 0 {
		return nil, false
	}
	resp := &RBACEnforceResp{}
	if resp, err := resp.Unmarshal(v); err != nil {
		log.Println("Unmarshal RbacContextAccount", err)
		return nil, false
	} else {
		return resp, true
	}
}
