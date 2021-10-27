package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/bkit/util/icopier"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

func (svc *Service) LoginAccount(ctx context.Context, in *model.LoginAccountReq, out *model.LoginAccountResp) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{
		AppId:    in.AppId,
		Username: in.Username,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrParamInvalid.MultiMsg("username or password invalid")
	}
	if acc.Status != types.LimitNormal {
		return errc.ErrAuthInvalid.MultiMsg("username locked")
	}

	if acc.LoginLock > time.Now().Unix() {
		return errc.ErrAuthInvalid.MultiMsg("username login locked")
	}

	if acc.Password != str.PasswordSlatMD5(in.Password, acc.Salt) {
		// 增加密码错误次数
		cols := []string{"pwd_wrong_num"}
		acc.PwdWrong += 1
		if acc.PwdWrong > 5 {
			cols = append(cols, "login_lock_at")
			acc.LoginLock = time.Now().AddDate(0, 0, 1).Unix()
		}
		if err := svc.d.UpdateAccount(ctx, acc, cols); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		return errc.ErrParamInvalid.MultiMsg("username or password invalid")
	}

	acc.GenToken()
	acc.GenTokenExpiredAt()
	acc.PwdWrong = 0
	acc.LoginLock = 0

	// 更新Token, 更新错误次数
	if err := svc.d.UpdateAccount(ctx, acc, []string{"token", "token_expired", "pwd_wrong", "login_lock"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	out.Token = acc.Token
	out.TokenExpired = acc.TokenExpired

	return nil
}

func (svc *Service) LoginOutAccount(ctx context.Context, in *model.LoginOutAccountReq) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{
		Token: in.Token,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return nil
	}
	acc.Token = ""
	acc.TokenExpired = 0

	if err := svc.d.UpdateAccount(ctx, acc, []string{"token", "token_expired"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) ListAccount(ctx context.Context, in *model.ListAccountReq, out *typ.ListResp) error {
	c, records, err := svc.d.ListAccount(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	list := make([]*model.ListAccount, 0, len(records))
	for _, v := range records {
		d := &model.ListAccount{
			Id:           v.Id,
			AppName:      "",
			AppId:        v.AppId,
			Nickname:     v.Nickname,
			Username:     v.Username,
			PwdWrong:     v.PwdWrong,
			LoginLock:    v.LoginLock,
			TokenExpired: v.TokenExpired,
			Memo:         v.Memo,
			Status:       v.Status,
			Roles:        make([]model.RoleBase, 0),
			UpdatedAt:    v.UpdatedAt.Unix(),
			CreatedAt:    v.CreatedAt.Unix(),
		}
		if exists, app, err := svc.d.GetAppConfigFromCache(ctx, &model.GetAppConfigReq{
			AppId: v.AppId,
		}); err == nil && exists {
			d.AppName = app.Name
		}
		roles := make([]model.RoleBase, 0)
		for _, rId := range v.Roles.Unmarshal() {
			if exists, role, err := svc.d.GetRoleConfigFromCache(ctx, &model.GetRoleConfigReq{
				Id: rId,
			}); err == nil && exists {
				roles = append(roles, model.RoleBase{
					Id:     role.Id,
					Name:   role.Name,
					Status: role.Status,
				})
			}
		}
		d.Roles = roles

		list = append(list, d)
	}

	out.Count = c
	out.List = list
	return nil
}

func (svc *Service) VerifyAccountToken(ctx context.Context, token string, out *model.VerifyAccountTokenResp) error {
	exists, acc, err := svc.d.GetAccountByTokenFromCache(ctx, token)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return fmt.Errorf("token not found")
	}
	if acc.TokenExpired < time.Now().Unix() {
		return fmt.Errorf("token expired")
	}
	if acc.LoginLock > time.Now().Unix() {
		return fmt.Errorf("account locked")
	}
	if err := icopier.Copy(out, acc); err != nil {
		return err
	}
	return nil
}

func (svc *Service) CreateAccount(ctx context.Context, in *model.CreateAccountReq) error {
	exists, _, err := svc.d.GetAccount(ctx, &model.GetAccountReq{AppId: in.AppId, Username: in.Username})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if exists {
		return errc.ErrParamInvalid.MultiMsg("username exists")
	}

	acc := &model.Account{
		AppId:        in.AppId,
		Nickname:     in.Nickname,
		Username:     in.Username,
		Salt:         str.RandAlphaNumString(6),
		PwdWrong:     0,
		LoginLock:    0,
		Token:        "",
		TokenExpired: 0,
		Memo:         "",
		Status:       types.LimitNormal,
		IsRoot:       0,
		Roles:        "",
	}
	acc.Password = str.PasswordSlatMD5(in.Password, acc.Salt)

	if err := svc.d.CreateAccount(ctx, acc); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAccountPassword(ctx context.Context, in *model.UpdateAccountPasswordReq) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Id: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("account")
	}
	acc.Password = str.PasswordSlatMD5(in.Password, acc.Salt)
	acc.Token = ""
	acc.LoginLock = 0
	acc.TokenExpired = 0
	acc.PwdWrong = 0

	if err := svc.d.UpdateAccount(ctx, acc, []string{"password", "token", "token_expired", "pwd_wrong", "login_lock"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAccountRole(ctx context.Context, in *model.UpdateAccountRoleReq) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Id: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("account")
	}

	for _, rId := range in.Roles {
		exists, _, err := svc.d.GetRoleConfigFromCache(ctx, &model.GetRoleConfigReq{
			Id: rId,
		})
		if err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		if !exists {
			return errc.ErrNotFound.MultiMsg(fmt.Sprintf("role id %d", rId))
		}
	}

	acc.Roles = new(types.IntSplitStr).Marshal(in.Roles)
	if err := svc.d.UpdateAccount(ctx, acc, []string{"roles"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	return nil
}

func (svc *Service) DelAccount(ctx context.Context, in *model.DelAccountReq) error {
	if err := svc.d.DelAccount(ctx, in.Id); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}
