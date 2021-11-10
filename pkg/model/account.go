package model

import (
	"fmt"
	"github.com/bbdshow/bkit/gen/str"
	"github.com/bbdshow/gin-rabc/pkg/types"
	"time"
)

type Account struct {
	Id        int64             `xorm:"not null pk autoincr BIGINT(20)"`
	Nickname  string            `xorm:"not null VARCHAR(64) comment('昵称')"`
	Username  string            `xorm:"not null VARCHAR(64) unique comment('账号名')"`
	Password  string            `xorm:"not null VARCHAR(64) comment('密码')"`
	Salt      string            `xorm:"not null VARCHAR(6) comment('盐')"`
	PwdWrong  int               `xorm:"not null TINYINT(4) comment('密码错误次数')"`
	LoginLock int64             `xorm:"not null BIGINT(20) comment('登录锁定时间')"`
	Memo      string            `xorm:"not null VARCHAR(128) comment('备注')"`
	Status    types.LimitStatus `xorm:"not null TINYINT(2) comment('状态 1-正常 2-锁定')"`
	UpdatedAt time.Time         `xorm:"updated"`
	CreatedAt time.Time         `xorm:"created"`
}

// AccountAppActivate 账户APP激活
type AccountAppActivate struct {
	Id           int64             `xorm:"not null pk autoincr BIGINT(20)"`
	AccountId    int64             `xorm:"not null BIGINT(20) unique(accountId_appId) comment('账户ID')"`
	AppId        string            `xorm:"not null VARCHAR(6) unique(accountId_appId) comment('APP分组')"`
	Token        string            `xorm:"not null VARCHAR(32) unique comment('Token')"`
	TokenExpired int64             `xorm:"not null BIGINT(20) comment('Token过期时间')"`
	Roles        types.IntSplitStr `xorm:"not null TEXT comment('角色ID')"`
	UpdatedAt    time.Time         `xorm:"updated"`
	CreatedAt    time.Time         `xorm:"created"`
}

func (a *AccountAppActivate) GenToken() string {
	return str.Md5String(str.RandAlphaNumString(12), fmt.Sprintf("%d_%s_%d", a.AccountId, a.AppId, time.Now().UnixNano()))
}

func (a *AccountAppActivate) GenTokenExpiredAt() int64 {
	return time.Now().AddDate(0, 0, 1).Unix()
}

//func (a *Account) GenToken() {
//	a.Token = str.Md5String(str.RandAlphaNumString(12), fmt.Sprintf("%d%s%d", a.Id, a.Username, time.Now().UnixNano()))
//}
//
//func (a *Account) GenTokenExpiredAt() {
//	a.TokenExpired = time.Now().AddDate(0, 0, 1).Unix()
//}
