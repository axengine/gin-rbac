package service

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/bkit/util/icopier"
	"time"
)

func (svc *Service) ListAppConfig(ctx context.Context, in *model.ListAppConfigReq, out *typ.ListResp) error {
	c, records, err := svc.d.ListAppConfig(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	list := make([]*model.ListAppConfig, 0, len(records))
	for _, v := range records {
		list = append(list, &model.ListAppConfig{
			Id:        v.Id,
			Name:      v.Name,
			AccessKey: v.AccessKey,
			SecretKey: v.SecretKey,
			Status:    v.Status,
			Memo:      v.Memo,
			UpdatedAt: v.UpdatedAt.Unix(),
			CreatedAt: v.CreatedAt.Unix(),
		})
	}
	out.Count = c
	out.List = list

	return nil
}

func (svc *Service) CreateAppConfig(ctx context.Context, in *model.CreateAppConfigReq) error {
	ak := str.RandAlphaNumString(16, true)
	time.Sleep(time.Millisecond)
	sk := str.RandAlphaNumString(32, true)
	d := &model.AppConfig{
		Name:      in.Name,
		AccessKey: ak,
		SecretKey: sk,
		Status:    1,
		Memo:      in.Memo,
	}
	if err := svc.d.CreateAppConfig(ctx, d); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) UpdateAppConfig(ctx context.Context, in *model.UpdateAppConfigReq) error {
	d := &model.AppConfig{
		Id: in.Id,
	}
	cols := make([]string, 0)
	if len(in.Name) > 0 {
		cols = append(cols, "name")
		d.Name = in.Name
	}
	if len(in.Memo) > 0 {
		cols = append(cols, "memo")
		d.Memo = in.Memo
	}
	if in.Status > 0 {
		cols = append(cols, "status")
		d.Status = in.Status
	}
	if in.SecretKey {
		cols = append(cols, "secret_key")
		d.SecretKey = str.RandAlphaNumString(32, true)
	}

	if err := svc.d.UpdateAppConfig(ctx, d, cols); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}

func (svc *Service) GetAppConfig(ctx context.Context, in *model.GetAppConfigReq, out *model.GetAppConfigResp) error {
	exists, c, err := svc.d.GetAppConfig(ctx, in)
	if err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	if !exists {
		return errc.ErrNotFound
	}

	if c.Status == types.LimitLocked {
		return errc.ErrAuthInvalid.MultiMsg("limit locked")
	}

	if err := icopier.Copy(out, c); err != nil {
		return errc.ErrInternalErr.MultiErr(err)
	}
	return nil
}
