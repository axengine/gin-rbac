package dao

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/errc"
	"xorm.io/xorm"
)

func (d *Dao) ListMenuConfig(ctx context.Context, in *model.ListMenuConfigReq) (int64, []*model.MenuConfig, error) {
	return 0, nil, nil
}

func (d *Dao) CreateMenuConfig(ctx context.Context, in *model.MenuConfig) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateMenuConfig(ctx context.Context, in *model.MenuConfig, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) UpsertMenuActionId() {}

func (d *Dao) UpsertMenuChildrenId() {}

func (d *Dao) ListActionConfig() {}

func (d *Dao) UpsertActionConfig(ctx context.Context, in *model.UpsertActionConfigReq) error {
	err := d.mysql.Transaction(func(sess *xorm.Session) error {
		ac := &model.ActionConfig{}
		exists, err := sess.Context(ctx).Where("app_id = ?", in.AppId).And("path = ?", in.Path).
			And("method = ?", in.Method).Get(ac)
		if err != nil {
			return errc.WithStack(err)
		}
		if !exists {
			r := &model.ActionConfig{
				AppId:  in.AppId,
				Name:   in.Name,
				Path:   in.Path,
				Method: in.Method,
			}
			if _, err := sess.Context(ctx).InsertOne(r); err != nil {
				return errc.WithStack(err)
			}
			return nil
		}

		if in.Name == ac.Name {
			return nil
		}
		ac.Name = in.Name
		if _, err := sess.Context(ctx).ID(ac.Id).Cols("name").Update(ac); err != nil {
			return errc.WithStack(err)
		}
		return nil
	})
	return errc.WithStack(err)
}

func (d *Dao) DelActionConfig() {}
