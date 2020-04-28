package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) PrivilegeMenuNew(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuNewReqProto) (*privilegeproto.PrivilegeMenuNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role menu new (%d, %d) err (recover %v).",
				reqproto.GetPrivilegeId(), reqproto.GetMenuId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("action new"), errMsg))
		}
	}()

	// todo : 校验权限

	roleMenuNewRespProto := &privilegeproto.PrivilegeMenuNewRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return roleMenuNewRespProto, nil
}
