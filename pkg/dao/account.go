package dao

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

func (d *Dao) ListAccount(ctx context.Context, in *model.ListAccountReq) (int64, []*model.Account, error) {
	sess := d.mysql.Context(ctx).Where("1 = 1")
	if len(in.Nickname) > 0 {
		sess.And("nickname like ?", "%"+in.Nickname+"%")
	}
	if len(in.Username) > 0 {
		sess.And("username like ?", "%"+in.Username+"%")
	}
	if len(in.AppId) > 0 {
		sess.And("app_id = ?", in.AppId)
	}
	if in.Status > 0 {
		sess.And("status = ?", in.Status)
	}

	records := make([]*model.Account, 0, in.Size)
	c, err := sess.OrderBy("id DESC").Limit(in.LimitStart()).FindAndCount(&records)
	return c, records, errc.WithStack(err)
}

func (d *Dao) GetAccount(ctx context.Context, in *model.GetAccountReq) (bool, *model.Account, error) {
	conds := make([]builder.Cond, 0)
	if in.Id > 0 {
		conds = append(conds, builder.Eq{"id": in.Id})
	}

	if in.Username != "" {
		conds = append(conds, builder.Like{"username", "%" + in.Username + "%"})
	}

	if len(conds) == 0 {
		return false, nil, errc.ErrParamInvalid.MultiMsg("condition required")
	}
	sess := d.mysql.Context(ctx).Where("1 = 1")
	for _, c := range conds {
		sess.And(c)
	}
	r := &model.Account{}
	exists, err := sess.Get(r)
	return exists, r, errc.WithStack(err)
}

func (d *Dao) GetAccountAppActivate(ctx context.Context, in *model.GetAccountAppActivateReq) (bool, *model.AccountAppActivate, error) {
	conds := make([]builder.Cond, 0)
	if in.Id > 0 {
		conds = append(conds, builder.Eq{"id": in.Id})
	}

	if in.AccountId > 0 {
		conds = append(conds, builder.Eq{"account_id": in.AccountId})
	}
	if in.AppId != "" {
		conds = append(conds, builder.Eq{"app_id": in.AppId})
	}

	if in.Token != "" {
		conds = append(conds, builder.Eq{"token": in.Token})
	}

	if len(conds) == 0 {
		return false, nil, errc.ErrParamInvalid.MultiMsg("condition required")
	}
	sess := d.mysql.Context(ctx).Where("1 = 1")
	for _, c := range conds {
		sess.And(c)
	}
	r := &model.AccountAppActivate{}
	exists, err := sess.Get(r)
	return exists, r, errc.WithStack(err)
}

func (d *Dao) CreateAccount(ctx context.Context, in *model.Account) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateAccount(ctx context.Context, in *model.Account, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) CreateAccountAppActivate(ctx context.Context, in *model.AccountAppActivate) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateAccountAppActivate(ctx context.Context, in *model.AccountAppActivate, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) UpsertAccountAppActivateRole(ctx context.Context, in []*model.AccountAppActivate) error {
	if len(in) <= 0 {
		return nil
	}
	err := d.mysql.Transaction(func(sess *xorm.Session) error {
		for _, v := range in {
			activate := &model.AccountAppActivate{}
			exists, err := sess.Context(ctx).Where("account_id = ?", v.AccountId).
				And("app_id = ?", v.AppId).Get(activate)
			if err != nil {
				return errc.WithStack(err)
			}
			if !exists {
				if _, err := sess.Context(ctx).InsertOne(v); err != nil {
					return errc.WithStack(err)
				}
			} else {
				activate.Roles = v.Roles
				if _, err := sess.Context(ctx).ID(activate.Id).Cols("roles").Update(activate); err != nil {
					return errc.WithStack(err)
				}
			}
		}
		return nil
	})
	return err
}

func (d *Dao) FindAccount(ctx context.Context, in *model.FindAccountReq) ([]*model.Account, error) {
	sess := d.mysql.Context(ctx).Where("1 = 1")
	if len(in.AppId) > 0 {
		sess.And("app_id = ?", in.AppId)
	}
	if in.Status > 0 {
		sess.And("status = ?", in.Status)
	}
	records := make([]*model.Account, 0)
	err := sess.Find(&records)
	return records, errc.WithStack(err)
}

func (d *Dao) FindAccountAppActivate(ctx context.Context, in *model.FindAccountAppActivateReq) ([]*model.AccountAppActivate, error) {
	sess := d.mysql.Context(ctx).Where("1 = 1")
	if in.AppId != "" {
		sess.And("app_id = ?", in.AppId)
	}
	if in.AccountId > 0 {
		sess.And("account_id = ?", in.AccountId)
	}
	records := make([]*model.AccountAppActivate, 0)
	err := sess.Find(&records)
	return records, errc.WithStack(err)
}
func (d *Dao) GetAccountFromCache(ctx context.Context, id int64) (bool, *model.Account, error) {
	key := fmt.Sprintf("Account_id_%d", id)
	v, err := d.memCache.Get(key)
	if err == nil {
		c, ok := v.(*model.Account)
		if ok {
			return true, c, nil
		}
	}
	exists, c, err := d.GetAccount(ctx, &model.GetAccountReq{Id: id})
	if err != nil {
		return false, nil, errc.WithStack(err)
	}
	if !exists {
		return false, nil, nil
	}
	_ = d.memCache.SetWithTTL(key, c, 15*time.Minute)
	return true, c, nil
}

func (d *Dao) GetAccountAppActivateFromCache(ctx context.Context, token string) (bool, *model.AccountAppActivate, error) {
	key := fmt.Sprintf("AccountAppActivate_token_%s", token)
	v, err := d.memCache.Get(key)
	if err == nil {
		c, ok := v.(*model.AccountAppActivate)
		if ok {
			return true, c, nil
		}
	}
	exists, c, err := d.GetAccountAppActivate(ctx, &model.GetAccountAppActivateReq{Token: token})
	if err != nil {
		return false, nil, errc.WithStack(err)
	}
	if !exists {
		return false, nil, nil
	}
	_ = d.memCache.SetWithTTL(key, c, 15*time.Minute)
	return true, c, nil
}

func (d *Dao) DelAccount(ctx context.Context, id int64) error {
	_, err := d.mysql.Context(ctx).ID(id).Delete(&model.Account{})
	return errc.WithStack(err)
}
