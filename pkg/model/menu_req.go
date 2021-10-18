package model

type ListMenuConfigReq struct {
}

type CreateMenuConfigReq struct {
	AppId    int64   `json:"appId"`
	Name     string  `json:"name"`
	Memo     string  `json:"memo"`
	ParentId int64   `json:"parentId"`
	Actions  []int64 `json:"actions"`
}

type UpsertActionConfigReq struct {
	AppId  int64  `json:"appId"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method int32  `json:"method"`
}
