package dao

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"time"
	"xorm.io/builder"
)

func (d *Dao) CreateAppConfig(ctx context.Context, in *model.AppConfig) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateAppConfig(ctx context.Context, in *model.AppConfig, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) ListAppConfig(ctx context.Context, in *model.ListAppConfigReq) (int64, []*model.AppConfig, error) {
	sess := d.mysql.Context(ctx).Where("1 = 1")
	if len(in.Name) > 0 {
		sess.And("name like ?", "%"+in.Name+"%")
	}

	if in.Status > 0 {
		sess.And("status = ?", in.Status)
	}

	records := make([]*model.AppConfig, 0, in.Size)
	c, err := sess.OrderBy("id DESC").Limit(in.LimitStart()).FindAndCount(&records)
	return c, records, errc.WithStack(err)
}

func (d *Dao) GetAppConfig(ctx context.Context, in *model.GetAppConfigReq) (bool, *model.AppConfig, error) {
	conds := make([]builder.Cond, 0)
	if len(in.AppId) > 0 {
		conds = append(conds, builder.Eq{"app_id": in.AppId})
	}
	if len(in.AccessKey) > 0 {
		conds = append(conds, builder.Eq{"access_key": in.AccessKey})
	}
	if len(conds) == 0 {
		return false, nil, errc.ErrParamInvalid.MultiMsg("condition required")
	}
	sess := d.mysql.Context(ctx).Where("1 = 1")
	for _, c := range conds {
		sess.And(c)
	}

	r := &model.AppConfig{}
	exists, err := sess.Get(r)
	return exists, r, errc.WithStack(err)
}

func (d *Dao) GetAppConfigFromCache(ctx context.Context, in *model.GetAppConfigReq) (bool, *model.AppConfig, error) {
	key := fmt.Sprintf("AppConfig_appId_%s_accessKey_%s", in.AppId, in.AccessKey)

	v, err := d.memCache.Get(key)
	if err == nil {
		c, ok := v.(*model.AppConfig)
		if ok {
			return true, c, nil
		}
	}
	exists, c, err := d.GetAppConfig(ctx, in)
	if err != nil {
		return false, nil, errc.WithStack(err)
	}
	if !exists {
		return false, nil, nil
	}
	_ = d.memCache.SetWithTTL(key, c, 5*time.Minute)
	return true, c, nil
}
