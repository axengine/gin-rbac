package service

import (
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/tests"
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

func TestService_GetAccountMenuAuth(t *testing.T) {
	out := &model.GetAccountMenuAuthResp{}
	if err := svc.GetAccountMenuAuth(ctx, &model.GetAccountMenuAuthReq{
		Token: "f77170d603b5f265e3198dafd2f6eba1",
	}, out); err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(out)
}
