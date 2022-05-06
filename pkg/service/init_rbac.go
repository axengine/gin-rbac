package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var swaggerJSON string

func init() {
	filename := "./docs/swagger.json"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Println("WARNING: ./docs/swagger.json not exists")
		return
	}
	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("WARNING: read ./docs/swagger.json err " + err.Error())
		return
	}
	swaggerJSON = string(byt)
}

func (svc *Service) InitRBAC(ctx context.Context) error {
	initAppId := "000000"
	exists, _, err := svc.d.GetAppConfig(ctx, &model.GetAppConfigReq{
		AppId: initAppId,
	})
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("already exists rbac config, can't init RBAC")
	}
	if swaggerJSON == "" {
		return fmt.Errorf("swagger JSON text, can't decode")
	}
	if err := svc.SwaggerJSONToActions(ctx, &model.SwaggerJSONToActionsReq{
		AppId:      initAppId,
		Prefix:     "/rbac/v1",
		SwaggerTxt: swaggerJSON,
	}); err != nil {
		return err
	}
	ak := str.RandAlphaNumString(16, true)
	time.Sleep(time.Millisecond)
	sk := str.RandAlphaNumString(32, true)
	d := &model.AppConfig{
		Name:      "RBAC",
		AppId:     initAppId,
		AccessKey: ak,
		SecretKey: sk,
		Status:    1,
		Memo:      "RBAC init root",
	}
	if err := svc.d.CreateAppConfig(ctx, d); err != nil {
		return err
	}
	actions, err := svc.d.FindActionConfig(ctx, &model.FindActionConfigReq{
		AppId: initAppId,
	})
	if err != nil {
		return err
	}
	actionsMap := map[string]*model.ActionConfig{}
	for _, v := range actions {
		actionsMap[v.Path] = v
	}
	accountActions := make([]string, 0)
	roleActions := make([]string, 0)
	menuActions := make([]string, 0)
	actionActions := make([]string, 0)
	appActions := make([]string, 0)

	for _, v := range accountMenus {
		action, ok := actionsMap[v]
		if !ok {
			return fmt.Errorf("path: %s not exists", v)
		}
		accountActions = append(accountActions, fmt.Sprintf("%d", action.Id))
	}

	for _, v := range roleMenus {
		action, ok := actionsMap[v]
		if !ok {
			return fmt.Errorf("path: %s not exists", v)
		}
		roleActions = append(roleActions, fmt.Sprintf("%d", action.Id))
	}
	for _, v := range menuMenus {
		action, ok := actionsMap[v]
		if !ok {
			return fmt.Errorf("path: %s not exists", v)
		}
		menuActions = append(menuActions, fmt.Sprintf("%d", action.Id))
	}
	for _, v := range actionMenus {
		action, ok := actionsMap[v]
		if !ok {
			return fmt.Errorf("path: %s not exists", v)
		}
		actionActions = append(actionActions, fmt.Sprintf("%d", action.Id))
	}
	for _, v := range appMenus {
		action, ok := actionsMap[v]
		if !ok {
			return fmt.Errorf("path: %s not exists", v)
		}
		appActions = append(appActions, fmt.Sprintf("%d", action.Id))
	}

	// 创建菜单
	if err := svc.d.CreateMenuConfig(ctx, &model.MenuConfig{
		AppId:    initAppId,
		Name:     "账户配置",
		Typ:      1,
		Memo:     "",
		ParentId: 0,
		Status:   1,
		Sequence: 0,
		Path:     "/accountset",
		Actions:  types.IntSplitStr(strings.Join(accountActions, ",")),
	}); err != nil {
		return err
	}
	if err := svc.d.CreateMenuConfig(ctx, &model.MenuConfig{
		AppId:    initAppId,
		Name:     "角色配置",
		Typ:      1,
		Memo:     "",
		ParentId: 0,
		Status:   1,
		Sequence: 0,
		Path:     "/roleset",
		Actions:  types.IntSplitStr(strings.Join(roleActions, ",")),
	}); err != nil {
		return err
	}
	if err := svc.d.CreateMenuConfig(ctx, &model.MenuConfig{
		AppId:    initAppId,
		Name:     "菜单配置",
		Typ:      1,
		Memo:     "",
		ParentId: 0,
		Status:   1,
		Sequence: 0,
		Path:     "/menuset",
		Actions:  types.IntSplitStr(strings.Join(menuActions, ",")),
	}); err != nil {
		return err
	}
	if err := svc.d.CreateMenuConfig(ctx, &model.MenuConfig{
		AppId:    initAppId,
		Name:     "功能配置",
		Typ:      1,
		Memo:     "",
		ParentId: 0,
		Status:   1,
		Sequence: 0,
		Path:     "/actionset",
		Actions:  types.IntSplitStr(strings.Join(actionActions, ",")),
	}); err != nil {
		return err
	}
	if err := svc.d.CreateMenuConfig(ctx, &model.MenuConfig{
		AppId:    initAppId,
		Name:     "APP配置",
		Typ:      1,
		Memo:     "",
		ParentId: 0,
		Status:   1,
		Sequence: 0,
		Path:     "/appset",
		Actions:  types.IntSplitStr(strings.Join(appActions, ",")),
	}); err != nil {
		return err
	}

	if err := svc.CreateRoleConfig(ctx, &model.CreateRoleConfigReq{
		AppId:  initAppId,
		Name:   "RBAC ROOT",
		IsRoot: 1,
		Memo:   "RBAC init root",
	}); err != nil {
		return err
	}

	if err := svc.CreateAccount(ctx, &model.CreateAccountReq{
		Nickname: "rbac root init",
		Username: "rbac_admin",
		Password: str.Md5String("rbac_admin_123"),
		Memo:     "rbac root init",
	}); err != nil {
		return err
	}

	return nil
}

var (
	accountMenus = []string{"/rbac/v1/account/list", "/rbac/v1/role/update", "/rbac/v1/account/create", "/rbac/v1/account/delete", "/rbac/v1/account/pwd/reset", "/rbac/v1/account/role/update"}
	roleMenus    = []string{"/rbac/v1/role/list", "/rbac/v1/role/update", "/rbac/v1/role/create", "/rbac/v1/role/action", "/rbac/v1/role/action/upsert", "/rbac/v1/role/delete"}
	menuMenus    = []string{"/rbac/v1/menu/action/update", "/rbac/v1/menu/update", "/rbac/v1/menu/tree", "/rbac/v1/menu/actions", "/rbac/v1/menu/create", "/rbac/v1/menu/delete"}
	actionMenus  = []string{"/rbac/v1/action/create", "/rbac/v1/action/import", "/rbac/v1/action/find", "/rbac/v1/action/list", "/rbac/v1/action/delete", "/rbac/v1/action/update"}
	appMenus     = []string{"/rbac/v1/app/list", "/rbac/v1/app/update", "/rbac/v1/app/create", "/rbac/v1/app/delete"}
)
