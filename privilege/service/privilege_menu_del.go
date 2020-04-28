package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *basicPrivilegeApiService) PrivilegeMenuDel(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuDelReqProto) (*privilegeproto.PrivilegeMenuDelRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilege menu del (%d) err (recover %v).",
				reqproto.GetId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilege menu del"), errMsg))
		}
	}()

	// todo : 校验权限

	privilegeMenuDelRespProto := &privilegeproto.PrivilegeMenuDelRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
	}

	return privilegeMenuDelRespProto, nil
}
