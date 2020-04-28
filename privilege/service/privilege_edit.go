package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) PrivilegeEdit(ctx context.Context, reqproto *privilegeproto.PrivilegeEditReqProto) (*privilegeproto.PrivilegeEditRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilege edit (%s, %s) err (recover %v).",
				reqproto.GetParentId(), reqproto.GetName(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilege edit"), errMsg))
		}
	}()

	// todo : 校验权限

	privilegeEditRespProto := &privilegeproto.PrivilegeEditRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return privilegeEditRespProto, nil
}
