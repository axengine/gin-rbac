package http

import (
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/ginutil"
	"github.com/gin-gonic/gin"
)

const (
	AccessTokenHeader = "X-Access-Token"
)

func MidAccessTokenVerify() gin.HandlerFunc {
	if svc == nil {
		panic("service nil")
	}

	return func(c *gin.Context) {
		if !cfg.Admin.AuthEnable {
			c.Next()
			return
		}
		token := c.GetHeader(AccessTokenHeader)
		if len(token) != 32 {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiMsg("X-Access-Token"))
			c.Abort()
			return
		}

		_, err := svc.VerifyAccountToken(c.Request.Context(), token)
		if err != nil {
			ginutil.RespErr(c, errc.ErrAuthInvalid.MultiErr(err))
			c.Abort()
			return
		}
		c.Set(AccessTokenHeader, token)
	}
}

func GetContextAccessToken(c *gin.Context) (string, error) {
	if !cfg.Admin.AuthEnable {
		return "", nil
	}
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
