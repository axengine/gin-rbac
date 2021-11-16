package service

import (
	"context"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/logs"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"go.uber.org/zap"
	"strconv"
)

func (svc *Service) RBACEnforce(ctx context.Context, in *model.RBACEnforceReq, out *model.RBACEnforceResp) error {
	vat := &model.VerifyAccountTokenResp{}
	if err := svc.VerifyAccountToken(ctx, in.AccessToken, vat); err != nil {
		return err
	}
	if !vat.Verify {
		out.Message = vat.Message
		return errc.ErrAuthInvalid.MultiMsg(out.Message)
	}

	out.Verify = true
	out.AppId = vat.AppId
	out.Nickname = vat.Nickname
	out.Username = vat.Username

	pass, err := svc.enforce.Enforce(strconv.FormatInt(vat.AccountId, 10), in.Path, in.Method)
	if err != nil {
		logs.Qezap.Error("RBACEnforce", zap.Any("in", in), zap.Any("accountId", vat.AccountId), zap.Error(err))
		return errc.ErrAuthInternalErr.MultiErr(err)
	}
	out.ActionPass = pass
	return nil
}

func (svc *Service) GetSecretKey(accessKey string) (string, error) {
	exists, app, err := svc.d.GetAppConfigFromCache(context.Background(), &model.GetAppConfigReq{
		AccessKey: accessKey,
	})
	if err != nil {
		return "", err
	}
	if !exists {
		return "", nil
	}
	if app.Status != types.LimitNormal {
		return "", nil
	}
	return app.SecretKey, nil
}
