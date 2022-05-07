package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/bkit/typ"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"strings"
	"time"
)

//func init() {
//	filename := "./docs/swagger.json"
//	_, err := os.Stat(filename)
//	if os.IsNotExist(err) {
//		log.Println("WARNING: ./docs/swagger.json not exists")
//		return
//	}
//	byt, err := ioutil.ReadFile(filename)
//	if err != nil {
//		log.Println("WARNING: read ./docs/swagger.json err " + err.Error())
//		return
//	}
//	swaggerJSON = string(byt)
//}

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
	//if swaggerJSON == "" {
	//	return fmt.Errorf("swagger JSON text, can't decode")
	//}
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

	// 绑定角色
	_, acc, err := svc.d.GetAccount(ctx, &model.GetAccountReq{
		Username: "rbac_admin",
	})
	if err != nil {
		return err
	}
	_, role, err := svc.d.GetRoleConfig(ctx, &model.GetRoleConfigReq{
		Name: "RBAC ROOT",
	})
	if err != nil {
		return err
	}

	if err := svc.UpdateAccountRole(ctx, &model.UpdateAccountRoleReq{
		IdReq: typ.IdReq{Id: acc.Id},
		Roles: []int64{role.Id},
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

var swaggerJSON = `{
    "swagger": "2.0",
    "info": {
        "description": "gin rbac manage API",
        "title": "gin rbac",
        "contact": {},
        "version": "1.0.0"
    },
    "host": "127.0.0.1:49000",
    "basePath": "/",
    "paths": {
        "/rbac/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出/修改密码"
                ],
                "summary": "[RBAC 登录]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.LoginAccountResp"
                        }
                    }
                }
            }
        },
        "/rbac/loginOut": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出/修改密码"
                ],
                "summary": "[RBAC 登出]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginOutAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/password/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出/修改密码"
                ],
                "summary": "[修改密码]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAccountPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置删除]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置列表]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListAccount"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/menu/auth": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户菜单权限]",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.GetAccountMenuAuthResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/pwd/reset": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置密码重置]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ResetAccountPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/role/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置角色更改]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAccountRoleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/account/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 账户配置"
                ],
                "summary": "[账户配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置删除]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/find": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据功能ID，查询基础信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置查询]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FindActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.FindActionConfigResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/import": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "导入Swagger JSON 文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置导入Swagger]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SwaggerJSONToActionsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置列表]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "method",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "path",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListActionConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/action/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 功能配置"
                ],
                "summary": "[功能配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateActionConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置删除]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置列表]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListAppConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/select": {
            "get": {
                "description": "用于Select组件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置筛选列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SelectAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.SelectAppConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/app/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC APP配置"
                ],
                "summary": "[APP配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAppConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/action/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置功能更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMenuConfigActionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/actions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置功能]",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "menuId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.GetMenuActionsResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateMenuConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置删除]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelMenuConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/tree": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置树]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "appId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.GetMenuTreeDirsResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/menu/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 菜单配置"
                ],
                "summary": "[菜单配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateMenuConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/action": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[获取角色菜单功能]",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "roleId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.GetRoleMenuActionResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/action/upsert": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置功能绑定]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpsertRoleMenuActionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置创建]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置删除]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置列表]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "appId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.ListRoleConfig"
                        }
                    }
                }
            }
        },
        "/rbac/v1/role/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 角色配置"
                ],
                "summary": "[角色配置更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateRoleConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/ginutil.BaseResp"
                        }
                    }
                }
            }
        },
        "/rbac/v1/verify/enforce": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "http请求验证是否拥有权限， 通过 accessToken 和要验证的 Path Method, 请求需签名",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC HTTP验证权限"
                ],
                "summary": "[HTTP验证权限]",
                "parameters": [
                    {
                        "type": "string",
                        "name": "accessToken",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "method",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/model.RBACEnforceResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ginutil.BaseResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "traceId": {
                    "type": "string"
                }
            }
        },
        "model.Action": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ActionBase": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.CreateAccountReq": {
            "type": "object",
            "required": [
                "nickname",
                "password",
                "username"
            ],
            "properties": {
                "memo": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.CreateActionConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "method",
                "name",
                "path"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "method": {
                    "description": "GET POST PUT DELETE",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "model.CreateAppConfigReq": {
            "type": "object",
            "required": [
                "memo",
                "name"
            ],
            "properties": {
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.CreateMenuConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "name",
                "path",
                "typ"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "model.CreateRoleConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "name"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.DelAccountReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.DelActionConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.DelAppConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.DelMenuConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.DelRoleConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.FindActionConfigReq": {
            "type": "object",
            "required": [
                "appId"
            ],
            "properties": {
                "actionId": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "appId": {
                    "type": "string"
                }
            }
        },
        "model.FindActionConfigResp": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ActionBase"
                    }
                }
            }
        },
        "model.GetAccountMenuAuthResp": {
            "type": "object",
            "properties": {
                "dirs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuTreeDir"
                    }
                },
                "isRoot": {
                    "type": "boolean"
                }
            }
        },
        "model.GetMenuActionsResp": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Action"
                    }
                }
            }
        },
        "model.GetMenuTreeDirsResp": {
            "type": "object",
            "properties": {
                "dirs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuTreeDir"
                    }
                }
            }
        },
        "model.GetRoleMenuActionResp": {
            "type": "object",
            "properties": {
                "menuActions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuAction"
                    }
                }
            }
        },
        "model.ListAccount": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "loginLock": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "pwdWrong": {
                    "type": "integer"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RoleBase"
                    }
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.ListActionConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.ListAppConfig": {
            "type": "object",
            "properties": {
                "accessKey": {
                    "description": "访问KEY",
                    "type": "string"
                },
                "appId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "description": "备注",
                    "type": "string"
                },
                "name": {
                    "description": "APP名",
                    "type": "string"
                },
                "secretKey": {
                    "description": "加密KEY",
                    "type": "string"
                },
                "status": {
                    "description": "状态 1-正常 2-限制",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.ListRoleConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "appName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        },
        "model.LoginAccountReq": {
            "type": "object",
            "required": [
                "appId",
                "password",
                "username"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.LoginAccountResp": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "tokenExpired": {
                    "type": "integer"
                }
            }
        },
        "model.LoginOutAccountReq": {
            "type": "object"
        },
        "model.MenuAction": {
            "type": "object",
            "required": [
                "menuId"
            ],
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "menuId": {
                    "type": "integer"
                }
            }
        },
        "model.MenuTreeDir": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "appId": {
                    "type": "string"
                },
                "children": {
                    "$ref": "#/definitions/model.MenuTreeDirs"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "model.MenuTreeDirs": {
            "type": "array",
            "items": {
                "$ref": "#/definitions/model.MenuTreeDir"
            }
        },
        "model.RBACEnforceResp": {
            "type": "object",
            "properties": {
                "actionPass": {
                    "description": "false-无权限",
                    "type": "boolean"
                },
                "appId": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verify": {
                    "type": "boolean"
                }
            }
        },
        "model.ResetAccountPasswordReq": {
            "type": "object",
            "required": [
                "id",
                "password"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.RoleBase": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "appName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.SelectAppConfig": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "description": "备注",
                    "type": "string"
                },
                "name": {
                    "description": "APP名",
                    "type": "string"
                }
            }
        },
        "model.SelectAppConfigReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "model.SwaggerJSONToActionsReq": {
            "type": "object",
            "required": [
                "appId"
            ],
            "properties": {
                "appId": {
                    "type": "string"
                },
                "prefix": {
                    "type": "string"
                },
                "swaggerTxt": {
                    "type": "string"
                }
            }
        },
        "model.UpdateAccountPasswordReq": {
            "type": "object",
            "required": [
                "newPassword",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "model.UpdateAccountReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateAccountRoleReq": {
            "type": "object",
            "required": [
                "id",
                "roles"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.UpdateActionConfigReq": {
            "type": "object",
            "required": [
                "id",
                "method",
                "name",
                "path",
                "status"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "method": {
                    "description": "GET POST PUT DELETE",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "status": {
                    "description": "1-正常 2-锁定",
                    "type": "integer"
                }
            }
        },
        "model.UpdateAppConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isSecretKey": {
                    "description": "1 = 重置加密KEY",
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "description": "状态 1-正常 2-限制",
                    "type": "integer"
                }
            }
        },
        "model.UpdateMenuConfigActionReq": {
            "type": "object",
            "required": [
                "menuId"
            ],
            "properties": {
                "actionId": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "menuId": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateMenuConfigReq": {
            "type": "object",
            "required": [
                "id",
                "status"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "status": {
                    "description": "1-正常 2-锁定",
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateRoleConfigReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isRoot": {
                    "type": "integer"
                },
                "memo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.UpsertRoleMenuActionReq": {
            "type": "object",
            "required": [
                "roleId"
            ],
            "properties": {
                "menuActions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MenuAction"
                    }
                },
                "roleId": {
                    "type": "integer"
                }
            }
        }
    }
}`
