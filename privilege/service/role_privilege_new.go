package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) RolePrivilegeNew(ctx context.Context, reqproto *privilegeproto.RolePrivilegeNewReqProto) (*privilegeproto.RolePrivilegeNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role privilege new (%d, %d) err (recover %v).",
				reqproto.GetPrivilegeId(), reqproto.GetRoleId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role privilege"), errMsg))
		}
	}()

	// todo : 校验权限

	rolePrivilegeNewRespProto := &privilegeproto.RolePrivilegeNewRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return rolePrivilegeNewRespProto, nil
}
