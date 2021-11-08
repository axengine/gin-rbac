package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bbdshow/bkit/errc"
	"github.com/bbdshow/bkit/tests"
	"github.com/bbdshow/gin-rabc/pkg/model"
	"github.com/bbdshow/gin-rabc/pkg/types"
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

	actions := make([]*model.ImportActionConfigReq, 0)
	for p, methods := range s.Paths {
		if in.Prefix != "" && !strings.HasPrefix(p, in.Prefix) {
			continue
		}
		for m, c := range methods {
			ac := &model.ImportActionConfigReq{
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
		if err := svc.ImportActionConfig(ctx, v); err != nil {
			return errc.ErrInternalErr.MultiErr(err)
		}
	}
	return nil
}
