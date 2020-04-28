package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) PrivilegeMenuGet(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuGetReqProto) (*privilegeproto.PrivilegeMenuGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilege menu get (%d) err (recover %v).",
				reqproto.GetPrivilegeId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilege menu get"), errMsg))
		}
	}()

	// todo : 校验权限

	privilegeMenuGetRespProto := &privilegeproto.PrivilegeMenuGetRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return privilegeMenuGetRespProto, nil
}
