package service

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"testing"
)

func TestService_UpsertActionConfig(t *testing.T) {
	if err := svc.UpsertActionConfig(ctx, &model.UpsertActionConfigReq{
		AppId:  1,
		Name:   "语言列表",
		Path:   "/v1/language/list",
		Method: 1,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestService_CreateMenuConfig(t *testing.T) {
	if err := svc.CreateMenuConfig(ctx, &model.CreateMenuConfigReq{
		AppId:    1,
		Name:     "系统设置",
		Memo:     "",
		ParentId: 0,
		Actions:  []int64{1},
	}); err != nil {
		t.Fatal(err)
	}
}
