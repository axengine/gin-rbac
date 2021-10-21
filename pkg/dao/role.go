package dao

import (
	"context"
	"fmt"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/errc"
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

func (d *Dao) ListRoleConfig(ctx context.Context, in *model.ListRoleConfigReq) (int64, []*model.RoleConfig, error) {
	sess := d.mysql.Context(ctx).Where("1 = 1")
	if len(in.Name) > 0 {
		sess.And("name like ?", "%"+in.Name+"%")
	}
	if len(in.AppId) > 0 {
		sess.And("app_id = ?", in.AppId)
	}
	records := make([]*model.RoleConfig, 0, in.Size)
	c, err := sess.OrderBy("id DESC").Limit(in.LimitStart()).FindAndCount(&records)
	return c, records, errc.WithStack(err)
}

func (d *Dao) GetRoleConfig(ctx context.Context, in *model.GetRoleConfigReq) (bool, *model.RoleConfig, error) {
	conds := make([]builder.Cond, 0)
	if in.Id > 0 {
		conds = append(conds, builder.Eq{"id": in.Id})
	}

	if len(conds) == 0 {
		return false, nil, errc.ErrParamInvalid.MultiMsg("condition required")
	}
	sess := d.mysql.Context(ctx).Where("1 = 1")
	for _, c := range conds {
		sess.And(c)
	}

	r := &model.RoleConfig{}
	exists, err := sess.Get(r)
	return exists, r, errc.WithStack(err)
}

func (d *Dao) CreateRoleConfig(ctx context.Context, in *model.RoleConfig) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateRoleConfig(ctx context.Context, in *model.RoleConfig, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateMenuActions() {}

func (d *Dao) DelRoleConfig() {}

func (d *Dao) GetRoleConfigFromCache(ctx context.Context, in *model.GetRoleConfigReq) (bool, *model.RoleConfig, error) {
	key := fmt.Sprintf("RoleConfig_id_%d", in.Id)
	v, err := d.memCache.Get(key)
	if err == nil {
		c, ok := v.(*model.RoleConfig)
		if ok {
			return true, c, nil
		}
	}
	exists, c, err := d.GetRoleConfig(ctx, in)
	if err != nil {
		return false, nil, errc.WithStack(err)
	}
	if !exists {
		return false, nil, nil
	}
	_ = d.memCache.SetWithTTL(key, c, 5*time.Minute)
	return true, c, nil
}

func (d *Dao) FindAllRoleMenuAction(ctx context.Context, roleId int64) ([]*model.RoleMenuAction, error) {
	records := make([]*model.RoleMenuAction, 0)
	err := d.mysql.Context(ctx).Where("role_id = ?", roleId).Find(&records)
	return records, errc.WithStack(err)
}

func (d *Dao) UpdateRoleMenuAction(ctx context.Context, add []*model.RoleMenuAction, del []int64) error {
	err := d.mysql.Transaction(func(sess *xorm.Session) error {
		if len(add) > 0 {
			_, err := sess.Context(ctx).Insert(add)
			if err != nil {
				return errc.WithStack(err)
			}
		}

		if len(del) > 0 {
			_, err := sess.Context(ctx).In("id", del).Delete(&model.RoleMenuAction{})
			if err != nil {
				return errc.WithStack(err)
			}
		}
		return nil
	})

	return errc.WithStack(err)
}
