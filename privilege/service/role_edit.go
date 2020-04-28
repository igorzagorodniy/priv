package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) RoleEdit(ctx context.Context, reqproto *privilegeproto.RoleEditReqProto) (*privilegeproto.RoleEditRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role edit (%s) err (recover %v).",
				reqproto.GetName(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role edit"), errMsg))
		}
	}()

	// todo : 校验权限

	roleEditRespProto := &privilegeproto.RoleEditRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return roleEditRespProto, nil
}
