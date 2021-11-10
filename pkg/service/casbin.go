package service

import (
	"context"
	"fmt"
	"github.com/bbdshow/gin-rabc/pkg/dao"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
	casbinM "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

var casbinModel casbinM.Model

func init() {
	casbinModel = casbinM.NewModel()
	casbinModel.AddDef("r", "r", "sub, obj, act")
	casbinModel.AddDef("p", "p", "sub, obj, act")
	casbinModel.AddDef("g", "g", "_, _")
	casbinModel.AddDef("e", "e", "some(where (p.eft == allow))")
	casbinModel.AddDef("m", "m", "g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)")
}

type CasbinAdapter struct {
	d *dao.Dao
}

func NewCasbinAdapter(d *dao.Dao) *CasbinAdapter {
	return &CasbinAdapter{d: d}
}

func (ca *CasbinAdapter) LoadPolicy(m casbinM.Model) error {
	if err := ca.loadRole(m); err != nil {
		return err
	}

	if err := ca.loadAccount(m); err != nil {
		return err
	}

	return nil
}

// 加载角色策略(p,roleId,actionPath,actionMethod)
func (ca *CasbinAdapter) loadRole(m casbinM.Model) error {
	roles, err := ca.d.FindAllRole(context.Background())
	if err != nil {
		return err
	}
	for _, role := range roles {
		for _, act := range role.Actions {
			if act.Status == types.LimitNormal && act.Path != "" && act.Name != "" {
				persist.LoadPolicyLine(fmt.Sprintf("p,%d,%s,%s", role.RoleId, act.Path, act.Method), m)
			}
		}
	}
	return nil
}

// 加载用户策略(g,accountId,roleId)
func (ca *CasbinAdapter) loadAccount(m casbinM.Model) error {
	accounts, err := ca.d.FindAccount(context.Background(), &model.FindAccountReq{Status: types.LimitNormal})
	if err != nil {
		return err
	}
	for _, v := range accounts {
		// 找到所有激活的
		activates, err := ca.d.FindAccountAppActivate(context.Background(), &model.FindAccountAppActivateReq{
			AccountId: v.Id,
		})
		if err != nil {
			return err
		}
		for _, act := range activates {
			roles := act.Roles.Unmarshal()
			for _, roleId := range roles {
				persist.LoadPolicyLine(fmt.Sprintf("g,%d,%d", v.Id, roleId), m)
			}
		}
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (ca *CasbinAdapter) SavePolicy(model casbinM.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (ca *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (ca *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (ca *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
