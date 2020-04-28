package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) RolePrivilegeDel(ctx context.Context, reqproto *privilegeproto.RolePrivilegeDelReqProto) (*privilegeproto.RolePrivilegeDelRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role privilege del (%d) err (recover %v).",
				reqproto.GetId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role privilege del"), errMsg))
		}
	}()

	// todo : 校验权限

	rolePrivilegeDelRespProto := &privilegeproto.RolePrivilegeDelRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return rolePrivilegeDelRespProto, nil
}
