package middleware

import (
	"context"
	"github.com/bbdshow/bkit/auth/sign"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

const (
	AccessTokenHeader = "X-Access-Token"
	ContextKeyAccount = "ContextKeyAccount"
)

type VerifyAccountToken func(ctx context.Context, token string, out *model.VerifyAccountTokenResp) error

func AccessTokenVerify(enable bool, verify VerifyAccountToken) gin.HandlerFunc {
	if verify == nil {
		panic("service nil")
	}
	return func(c *gin.Context) {
		if !enable {
			c.Next()
			return
		}
		token := c.GetHeader(AccessTokenHeader)
		if len(token) != 32 {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiMsg("X-Access-Token"))
			c.Abort()
			return
		}
		vat := model.VerifyAccountTokenResp{}
		err := verify(c.Request.Context(), token, &vat)
		if err != nil {
			ginutil.RespErr(c, err)
			c.Abort()
			return
		}
		if !vat.Verify {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiMsg(vat.Message))
			c.Abort()
			return
		}
		c.Set(AccessTokenHeader, token)
		c.Set(ContextKeyAccount, vat)
	}
}

func GetContextAccessToken(c *gin.Context) (string, error) {
	token, exists := c.Get(AccessTokenHeader)
	if !exists {
		return "", errc.ErrNotFound.MultiMsg("context token")
	}
	t, ok := token.(string)
	if !ok {
		return "", errc.ErrAuthInternalErr.MultiMsg("token not string")
	}
	return t, nil
}

// RBACEnforce  the handler process should come after AccessTokenVerify
func RBACEnforce(enable bool, enforce *casbin.SyncedEnforcer, skipPrefixPaths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !enable {
			c.Next()
			return
		}
		path := c.FullPath()
		method := strings.ToUpper(c.Request.Method)
		// skip
		for _, p := range skipPrefixPaths {
			if strings.HasPrefix(path, p) {
				c.Next()
				return
			}
		}
		acc, exists := c.Get(ContextKeyAccount)
		if !exists {
			ginutil.RespErr(c, errc.ErrAuthRequired.MultiMsg("context account not found"))
			c.Abort()
			return
		}
		vat, ok := acc.(model.VerifyAccountTokenResp)
		if !ok {
			logs.Qezap.Error("MidRBACEnforce", zap.Any("account", vat))
			ginutil.RespErr(c, errc.ErrAuthInternalErr.MultiMsg("account invalid"))
			c.Abort()
			return
		}
		pass, err := enforce.Enforce(strconv.FormatInt(vat.AccountId, 10), path, method)
		if err != nil {
			logs.Qezap.Error("MidRBACEnforce", zap.Any("Enforce", err))
			ginutil.RespErr(c, errc.ErrInternalErr.MultiErr(err))
			c.Abort()
			return
		}
		if !pass {
			ginutil.RespErr(c, errc.NewError(15, "no access").MultiMsg("account not auth"))
			c.Abort()
			return
		}
		c.Next()
	}
}

type GetSecretKey func(accessKey string) (string, error)

func ApiSign(enable bool, secretKey GetSecretKey) gin.HandlerFunc {
	// API 签名中间件
	signConfig := ginutil.SignConfig{
		Enable: enable,
		Config: sign.Config{
			SignValidDuration: time.Minute,
			Method:            sign.HmacSha256Hex,
			PathSign:          false,
		},
		SupportMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}
	return ginutil.ApiSignVerify(&signConfig, secretKey)
}
