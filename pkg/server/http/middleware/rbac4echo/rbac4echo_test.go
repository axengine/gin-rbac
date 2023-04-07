package rbac4echo

import (
	"bytes"
	"encoding/json"
	"github.com/axengine/utils/hash"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func loginRBAC() (string, error) {
	type LoginAccountReq struct {
		AppId    string `json:"appId" binding:"required,len=6"`
		Username string `json:"username" binding:"required,lte=64"`
		Password string `json:"password" binding:"required,len=32"`
	}

	type LoginAccountResp struct {
		Token        string `json:"token"`
		TokenExpired int64  `json:"tokenExpired"`
		Nickname     string `json:"nickname"`
	}
	type DataResp struct {
		Code    int              `json:"code" xml:"code"`
		Message string           `json:"message" xml:"message"`
		Data    LoginAccountResp `json:"data"`
	}

	req := LoginAccountReq{
		AppId:    "OYYZRY",
		Username: "juncawallet",
		Password: hash.MD5Lower([]byte("000000")),
	}
	bz, _ := json.Marshal(&req)

	resp, err := http.Post("http://127.0.0.7:49000/rbac/login", "application/json", bytes.NewReader(bz))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var res DataResp
	bz, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(bz, &res); err != nil {
		return "", err
	}
	if res.Code != 0 {
		return "", errors.New(res.Message)
	}
	return res.Data.Token, nil
}
