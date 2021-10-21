package service

import (
	"github.com/bbdshow/admin-rabc/pkg/model"
	"testing"
)

func TestService_SwaggerJSONToActions(t *testing.T) {
	in := &model.SwaggerJSONToActionsReq{
		AppId: "000000",
		SwaggerTxt: `{
    "swagger": "2.0",
    "info": {
        "description": "admin rbac manage API",
        "title": "admin rbac",
        "contact": {},
        "version": "1.0.0"
    },
    "host": "API_HOST:49000",
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
                    "RBAC 登录/登出"
                ],
                "summary": "RBAC 登录",
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC 登录/登出"
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
        "/rbac/v1/account/list": {
            "get": {
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
                "summary": "[账户配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListAccountReq"
                        }
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
        "/rbac/v1/account/pwd/update": {
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
                "summary": "[账户配置密码更改]",
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
        "/rbac/v1/action/list": {
            "get": {
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
                "summary": "[功能配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListActionConfigReq"
                        }
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
        "/rbac/v1/action/upsert": {
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
                "summary": "[功能配置创建/更新]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpsertActionConfigReq"
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
        "/rbac/v1/app/list": {
            "get": {
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
                "summary": "[APP配置列表]",
                "parameters": [
                    {
                        "description": "request param",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListAppConfigReq"
                        }
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
        "model.CreateAccountReq": {
            "type": "object",
            "required": [
                "appId",
                "nickname",
                "pwd",
                "username"
            ],
            "properties": {
                "appId": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "pwd": {
                    "type": "string"
                },
                "username": {
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
        "model.ListAccount": {
            "type": "object",
            "properties": {
                "appId": {
                    "type": "integer"
                },
                "appName": {
                    "type": "string"
                },
                "createdAt": {
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
                        "type": "string"
                    }
                },
                "status": {
                    "type": "integer"
                },
                "tokenExpired": {
                    "type": "integer"
                }
            }
        },
        "model.ListAccountReq": {
            "type": "object",
            "required": [
                "page",
                "size"
            ],
            "properties": {
                "appId": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.ListActionConfig": {
            "type": "object"
        },
        "model.ListActionConfigReq": {
            "type": "object",
            "required": [
                "id",
                "page",
                "size"
            ],
            "properties": {
                "appId": {
                    "type": "integer"
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
                "page": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "size": {
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
        "model.ListAppConfigReq": {
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
                },
                "status": {
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
                    "type": "integer"
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
                "token": {
                    "type": "string"
                },
                "tokenExpired": {
                    "type": "integer"
                }
            }
        },
        "model.LoginOutAccountReq": {
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
        "model.UpdateAccountPasswordReq": {
            "type": "object",
            "required": [
                "id",
                "pwd"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "pwd": {
                    "type": "string"
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
        "model.UpdateAppConfigReq": {
            "type": "object",
            "required": [
                "id",
                "memo",
                "name"
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
                "secretKey": {
                    "description": "true = 重置加密KEY",
                    "type": "boolean"
                },
                "status": {
                    "description": "状态 1-正常 2-限制",
                    "type": "integer"
                }
            }
        },
        "model.UpsertActionConfigReq": {
            "type": "object",
            "required": [
                "appId",
                "method",
                "name",
                "path"
            ],
            "properties": {
                "appId": {
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
                }
            }
        }
    }
}`,
	}
	if err := svc.SwaggerJSONToActions(ctx, in); err != nil {
		t.Fatal(err)
	}
}