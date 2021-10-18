package service

import (
	"context"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"testing"
)

func TestService_CreateAppConfig(t *testing.T) {
	if err := svc.CreateAppConfig(context.Background(), &model.CreateAppConfigReq{
		Name: "Test",
		Memo: "Test",
	}); err != nil {
		t.Fatal(err)
	}
}
