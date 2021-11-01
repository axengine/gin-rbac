package model

type RBACEnforceReq struct {
	AccessToken string `json:"accessToken" binding:"required,len=32"`
	Path        string `json:"path" binding:"required"`
	Method      string `json:"method" binding:"required"`
}

type RBACEnforceResp struct {
	AccountId  int64  `json:"accountId"`
	AppId      string `json:"appId"`
	Nickname   string `json:"nickname"`
	IsRoot     int32  `json:"isRoot"`
	ActionPass bool   `json:"actionPass"` // false-无权限
}
