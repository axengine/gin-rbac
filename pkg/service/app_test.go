package service

import (
	"context"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"testing"
)

func TestService_CreateAppConfig(t *testing.T) {
	if err := svc.CreateAppConfig(context.Background(), &model.CreateAppConfigReq{
		Name: "RBAC",
		Memo: "RBAC itself",
	}); err != nil {
		t.Fatal(err)
	}
}
