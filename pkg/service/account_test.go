package service

import (
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	if err := svc.CreateAccount(ctx, &model.CreateAccountReq{
		AppId:    "000000",
		Nickname: "rbac root",
		Username: "rbacRoot",
		Password: str.Md5String("111111"),
	}); err != nil {
		t.Fatal(err)
	}
}
