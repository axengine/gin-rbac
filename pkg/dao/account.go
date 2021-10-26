package dao

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/errc"
	"xorm.io/builder"
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
	if len(in.AppId) > 0 {
		conds = append(conds, builder.Eq{"app_id": in.AppId})
	}

	if len(in.Username) > 0 {
		conds = append(conds, builder.Eq{"username": in.Username})
	}
	if len(in.Token) > 0 {
		conds = append(conds, builder.Eq{"token": in.Token})
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

func (d *Dao) CreateAccount(ctx context.Context, in *model.Account) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateAccount(ctx context.Context, in *model.Account, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
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

func (d *Dao) DelAccount() {}
