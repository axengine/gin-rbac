package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bbdshow/admin-rabc/pkg/model"
	"github.com/bbdshow/admin-rabc/pkg/types"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/tests"
	"path"
	"strings"
)

// SwaggerJSONToActions swagger.json 文件解析成Actions存储
func (svc *Service) SwaggerJSONToActions(ctx context.Context, in *model.SwaggerJSONToActionsReq) error {
	s := &model.SwaggerJSON{}
	if err := json.Unmarshal([]byte(in.SwaggerTxt), s); err != nil {
		return errc.ErrParamInvalid.MultiMsg(fmt.Sprintf("swaggerTxt %v", err))
	}
	tests.PrintBeautifyJSON(s)

	actions := make([]*model.UpsertActionConfigReq, 0)
	for p, methods := range s.Paths {
		for m, c := range methods {
			ac := &model.UpsertActionConfigReq{
				AppId:  in.AppId,
				Name:   c.Summary,
				Path:   path.Join(s.BasePath, p),
				Method: strings.ToUpper(m),
				Status: types.LimitNormal,
			}
			actions = append(actions, ac)
		}
	}
	for _, v := range actions {
		if err := svc.UpsertActionConfig(ctx, v); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
	}
	return nil
}
