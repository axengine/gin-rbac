package service

import (
	"github.com/bbdshow/bkit/tests"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"testing"
)

func TestService_RBACEnforce(t *testing.T) {
	out := &model.RBACEnforceResp{}
	if err := svc.RBACEnforce(ctx, &model.RBACEnforceReq{
		AccessToken: "d97ac6546c37054d7fa2f58d933400f0",
		Path:        "/admin/v1/chain/list",
		Method:      "GET",
	}, out); err != nil {
		t.Fatal(err)
	}
	tests.PrintBeautifyJSON(out)
}
