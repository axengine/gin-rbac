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
		Token: "2fd698f38ab6165ea0aee6f1ee646c55",
	}, out); err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(out)
}
