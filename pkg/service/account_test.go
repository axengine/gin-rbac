package service

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/bkit/gen/str"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	if err := svc.CreateAccount(ctx, &model.CreateAccountReq{
		AppId:    "000000",
		Nickname: "rbacRoot",
		Username: "rbacRoot",
		Password: str.Md5String("111111"),
	}); err != nil {
		t.Fatal(err)
	}
}
