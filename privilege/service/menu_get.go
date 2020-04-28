package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) MenuGet(ctx context.Context, reqproto *privilegeproto.MenuGetReqProto) (*privilegeproto.MenuGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("menu get (%s, %s) err (recover %v).",
				reqproto.GetName(), reqproto.GetRouter(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("menu get"), errMsg))
		}
	}()

	// todo : 校验权限

	menuGetRespProto := &privilegeproto.MenuGetRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return menuGetRespProto, nil
}
