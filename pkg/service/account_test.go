package service

import (
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/tests"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	if err := svc.CreateAccount(ctx, &model.CreateAccountReq{
		Nickname: "rbac root",
		Username: "rbacRoot",
		Password: str.Md5String("111111"),
	}); err != nil {
		t.Fatal(err)
	}
}

func TestService_GetAccountMenuAuth(t *testing.T) {
	out := &model.GetAccountMenuAuthResp{}
	if err := svc.GetAccountMenuAuth(ctx, &model.GetAccountMenuAuthReq{
		Token: "fd7ee4391d4d1653b4ffde2b93459b5c",
	}, out); err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(out)
}
