package middleware

import (
	"context"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
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
		acc := model.VerifyAccountTokenResp{}
		err := verify(c.Request.Context(), token, &acc)
		if err != nil {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiErr(err))
			c.Abort()
			return
		}
		c.Set(AccessTokenHeader, token)
		c.Set(ContextKeyAccount, acc)
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
		account, ok := acc.(model.Account)
		if !ok {
			logs.Qezap.Error("MidRBACEnforce", zap.Any("account", account))
			ginutil.RespErr(c, errc.ErrAuthInternalErr.MultiMsg("account invalid"))
			c.Abort()
			return
		}
		if account.IsRoot == 1 {
			c.Next()
			return
		}
		pass, err := enforce.Enforce(account.Id, path, method)
		if err != nil {
			logs.Qezap.Error("MidRBACEnforce", zap.Any("Enforce", err))
			ginutil.RespErr(c, errc.ErrInternalErr.MultiErr(err))
			c.Abort()
			return
		}
		if !pass {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiMsg("account not auth"))
			c.Abort()
			return
		}
		c.Next()
	}
}
