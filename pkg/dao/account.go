package dao

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/errc"
)

func (d *Dao) ListAccount(ctx context.Context, in *model.ListAccountReq) (int64, []*model.Account, error) {
	return 0, nil, nil
}

func (d *Dao) CreateAccount(ctx context.Context, in *model.Account) error {
	_, err := d.mysql.Context(ctx).InsertOne(in)
	return errc.WithStack(err)
}

func (d *Dao) UpdateAccount(ctx context.Context, in *model.Account, cols []string) error {
	_, err := d.mysql.Context(ctx).ID(in.Id).Cols(cols...).Update(in)
	return errc.WithStack(err)
}

func (d *Dao) DelAccount() {}
