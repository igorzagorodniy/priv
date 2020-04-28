package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) MenuEdit(ctx context.Context, reqproto *privilegeproto.MenuEditReqProto) (*privilegeproto.MenuEditRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("menu edit (%s, %s) err (recover %v).",
				reqproto.GetName(), reqproto.GetRouter(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("menu edit"), errMsg))
		}
	}()

	// todo : 校验权限

	roleEditRespProto := &privilegeproto.MenuEditRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return roleEditRespProto, nil
}
