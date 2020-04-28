package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) MenuNew(ctx context.Context, reqproto *privilegeproto.MenuNewReqProto) (*privilegeproto.MenuNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("menu new (%s, %s) err (recover %v).",
				reqproto.GetName(), reqproto.GetRouter(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("menu new"), errMsg))
		}
	}()

	// todo : 校验权限

	menuNewRespProto := &privilegeproto.MenuNewRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return menuNewRespProto, nil
}
