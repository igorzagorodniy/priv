package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) RoleGet(ctx context.Context, reqproto *privilegeproto.RoleGetReqProto) (*privilegeproto.RoleGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role get (%s, %s) err (recover %v).",
				reqproto.GetName(), reqproto.GetFrom(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role get"), errMsg))
		}
	}()

	// todo : 校验权限

	roleGetRespProto := &privilegeproto.RoleGetRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return roleGetRespProto, nil
}
