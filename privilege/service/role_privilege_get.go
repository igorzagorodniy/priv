package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) RolePrivilegeGet(ctx context.Context, reqproto *privilegeproto.RolePrivilegeGetReqProto) (*privilegeproto.RolePrivilegeGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role privilege  get (%d, %d) err (recover %v).",
				reqproto.GetPrivilegeId(), reqproto.GetRoleId(), logutil.PrintPanic())
			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role privilege get"), errMsg))
		}
	}()

	// todo : 校验权限

	rolePrivilegeGetRespProto := &privilegeproto.RolePrivilegeGetRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}
	return rolePrivilegeGetRespProto, nil
}
