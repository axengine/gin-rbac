package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

func (svc *Service) LoginAccount(ctx context.Context, in *model.LoginAccountReq, out *model.LoginAccountResp) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{
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

	exists, activate, err := svc.d.GetAccountAppActivate(ctx, &model.GetAccountAppActivateReq{
		AccountId: acc.Id,
		AppId:     in.AppId,
	})
	if !exists {
		return errc.ErrAuthInvalid.MultiMsg("The account is not registered for this APP")
	}
	acc.PwdWrong = 0
	acc.LoginLock = 0

	// 更新错误次数
	if err := svc.d.UpdateAccount(ctx, acc, []string{"token", "token_expired", "pwd_wrong", "login_lock"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	// 更新Token
	activate.Token = activate.GenToken()
	activate.TokenExpired = activate.GenTokenExpiredAt()

	// 更新Token
	if err := svc.d.UpdateAccountAppActivate(ctx, activate, []string{"token", "token_expired"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	out.Token = activate.Token
	out.TokenExpired = activate.TokenExpired
	out.Nickname = acc.Nickname

	return nil
}

func (svc *Service) LoginOutAccount(ctx context.Context, in *model.LoginOutAccountReq) error {
	exists, activate, err := svc.d.GetAccountAppActivate(ctx, &model.GetAccountAppActivateReq{
		Token: in.Token,
	})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return nil
	}
	activate.Token = activate.GenToken()
	activate.TokenExpired = 0

	if err := svc.d.UpdateAccountAppActivate(ctx, activate, []string{"token", "token_expired"}); err != nil {
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
			Id:        v.Id,
			Nickname:  v.Nickname,
			Username:  v.Username,
			PwdWrong:  v.PwdWrong,
			LoginLock: v.LoginLock,
			Memo:      v.Memo,
			Status:    v.Status,
			Roles:     make([]model.RoleBase, 0),
			UpdatedAt: v.UpdatedAt.Unix(),
			CreatedAt: v.CreatedAt.Unix(),
		}
		// 查看激活的渠道
		activates, err := svc.d.FindAccountAppActivate(ctx, &model.FindAccountAppActivateReq{
			AccountId: v.Id,
		})
		if err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		roles := make([]model.RoleBase, 0)
		for _, act := range activates {
			for _, rId := range act.Roles.Unmarshal() {
				r := model.RoleBase{}
				if exists, role, err := svc.d.GetRoleConfigFromCache(ctx, &model.GetRoleConfigReq{
					Id: rId,
				}); err == nil && exists {
					r.Id = role.Id
					r.Name = role.Name
					r.Status = role.Status
					r.AppId = role.AppId
				}
				if r.AppId != "" {
					if exists, app, err := svc.d.GetAppConfigFromCache(ctx, &model.GetAppConfigReq{
						AppId: r.AppId,
					}); err == nil && exists {
						r.AppName = app.Name
					}
				}
				roles = append(roles, r)
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
	exists, activate, err := svc.d.GetAccountAppActivateFromCache(ctx, token)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		out.Message = "token not found"
		return nil
	}
	if activate.TokenExpired < time.Now().Unix() {
		out.Message = "token expired"
		return nil
	}

	exists, acc, err := svc.d.GetAccountFromCache(ctx, activate.AccountId)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		out.Message = "account not found"
		return nil
	}
	if acc.LoginLock > time.Now().Unix() {
		out.Message = "account locked"
		return nil
	}

	out.Verify = true
	out.AccountId = acc.Id
	out.Nickname = acc.Nickname
	out.Username = acc.Username
	out.AppId = activate.AppId
	return nil
}

func (svc *Service) CreateAccount(ctx context.Context, in *model.CreateAccountReq) error {
	exists, _, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Username: in.Username})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if exists {
		return errc.ErrParamInvalid.MultiMsg("username exists")
	}

	acc := &model.Account{
		Nickname: in.Nickname,
		Username: in.Username,
		Salt:     str.RandAlphaNumString(6),
		Memo:     in.Memo,
		Status:   types.LimitNormal,
	}
	acc.Password = str.PasswordSlatMD5(in.Password, acc.Salt)

	if err := svc.d.CreateAccount(ctx, acc); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAccount(ctx context.Context, in *model.UpdateAccountReq) error {
	r := &model.Account{
		Id:       in.Id,
		Nickname: in.Nickname,
		Memo:     in.Memo,
		Status:   in.Status,
	}
	cols := make([]string, 0)
	if in.Nickname != "" {
		cols = append(cols, "nickname")
	}
	if in.Memo != "" {
		cols = append(cols, "memo")
	}
	if in.Status > 0 {
		cols = append(cols, "status")
	}

	if err := svc.d.UpdateAccount(ctx, r, cols); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) ResetAccountPassword(ctx context.Context, in *model.ResetAccountPasswordReq) error {
	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Id: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("account")
	}
	acc.Password = str.PasswordSlatMD5(in.Password, acc.Salt)
	acc.LoginLock = 0
	acc.PwdWrong = 0
	if err := svc.d.UpdateAccount(ctx, acc, []string{"password", "pwd_wrong", "login_lock"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAccountPassword(ctx context.Context, in *model.UpdateAccountPasswordReq) error {
	exists, activate, err := svc.d.GetAccountAppActivate(ctx, &model.GetAccountAppActivateReq{Token: in.Token})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrAuthInvalid.MultiMsg("token")
	}

	exists, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Id: activate.AccountId})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("account")
	}

	if acc.Password != str.PasswordSlatMD5(in.OldPassword, acc.Salt) {
		return errc.ErrAuthInvalid.MultiMsg("old password")
	}

	acc.Password = str.PasswordSlatMD5(in.NewPassword, acc.Salt)
	acc.LoginLock = 0
	acc.PwdWrong = 0
	if err := svc.d.UpdateAccount(ctx, acc, []string{"password", "pwd_wrong", "login_lock"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	activate.TokenExpired = 0
	if err := svc.d.UpdateAccountAppActivate(ctx, activate, []string{"token_expired"}); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAccountRole(ctx context.Context, in *model.UpdateAccountRoleReq) error {
	exists, _, err := svc.d.GetAccount(ctx, &model.GetAccountReq{Id: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound.MultiMsg("account")
	}

	activates, err := svc.d.FindAccountAppActivate(ctx, &model.FindAccountAppActivateReq{AccountId: in.Id})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	type isChange struct {
		IsChange bool
		Value    *model.AccountAppActivate
	}
	activatesMap := map[string]*isChange{}
	for _, v := range activates {
		activatesMap[v.AppId] = &isChange{
			IsChange: false,
			Value:    v,
		}
	}

	for _, rId := range in.Roles {
		exists, role, err := svc.d.GetRoleConfigFromCache(ctx, &model.GetRoleConfigReq{
			Id: rId,
		})
		if err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		if !exists {
			return errc.ErrNotFound.MultiMsg(fmt.Sprintf("role id %d", rId))
		}
		// 如果没有激活APP，添加了角色就自动激活此APP
		act, ok := activatesMap[role.AppId]
		if !ok {
			a := &model.AccountAppActivate{
				AccountId: in.Id,
				AppId:     role.AppId,
				Roles:     new(types.IntSplitStr).Marshal([]int64{role.Id}),
			}
			a.Token = a.GenToken()
			activatesMap[role.AppId] = &isChange{
				IsChange: true,
				Value:    a,
			}
		} else {
			act.IsChange, act.Value.Roles = act.Value.Roles.Set(role.Id)
		}
	}

	ups := make([]*model.AccountAppActivate, 0)
	for _, v := range activatesMap {
		if v.IsChange {
			ups = append(ups, v.Value)
		}
	}
	if err := svc.d.UpsertAccountAppActivateRole(ctx, ups); err != nil {
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

func (svc *Service) GetAccountMenuAuth(ctx context.Context, in *model.GetAccountMenuAuthReq, out *model.GetAccountMenuAuthResp) error {
	exists, activate, err := svc.d.GetAccountAppActivate(ctx, &model.GetAccountAppActivateReq{Token: in.Token})
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrAuthInvalid.MultiMsg("token")
	}
	out.IsRoot = false

	dirs := make(model.MenuTreeDirs, 0)
	roles, err := svc.d.FindRoleConfig(ctx, &model.FindRoleConfigReq{RoleId: activate.Roles.Unmarshal()})
	if len(roles) <= 0 {
		out.Dirs = dirs
		return nil
	}
	isRoot := false
	roleId := make([]int64, 0, len(roles))
	for _, v := range roles {
		if v.IsRoot == 1 {
			isRoot = true
		}
		roleId = append(roleId, v.Id)
	}
	if isRoot {
		out.IsRoot = isRoot
		dir := &model.GetMenuTreeDirsResp{}
		if err := svc.GetMenuTreeDirs(ctx, &model.GetMenuTreeDirsReq{AppId: activate.AppId}, dir); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
		out.Dirs = dir.Dirs
		return nil
	}
	// 获取所有角色的菜单
	menuId, err := svc.d.GroupRolesMenuId(ctx, roleId)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	root, children, err := svc.d.FindMenuConfigAsRootOrChildren(ctx, menuId)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}

	dirs, err = svc.menuTreeDirs(append(root, children...))
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	out.IsRoot = false
	out.Dirs = dirs

	return nil
}
