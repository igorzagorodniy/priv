package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) ActionEdit(ctx context.Context, reqproto *privilegeproto.ActionEditReqProto) (*privilegeproto.ActionEditRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("action new (%s, %s) err (recover %v).",
				reqproto.GetActionName(), reqproto.GetActionPath(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("action new"), errMsg))
		}
	}()

	// todo : 校验权限

	actionEditRespProto := &privilegeproto.ActionEditRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return actionEditRespProto, nil
}
