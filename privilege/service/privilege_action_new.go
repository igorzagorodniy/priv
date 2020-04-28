package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
	privilegecligrpc "xtech-kit/svc/privilege-svc/client/grpc"
)

func (p *basicPrivilegeApiService) PrivilegeActionNew(ctx context.Context, reqproto *privilegeproto.PrivilegeActionNewReqProto) (*privilegeproto.PrivilegeActionNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilegeaction new (%v) err (recover %v).",
				reqproto, logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilegeaction new"), errMsg))
		}
	}()

	// TODO:业务逻辑判断

	privilegeactionCreateReqProto := privilegecligrpc.CreatePrivilegeActionCreateReqProto(
		reqproto.GetPrivilegeId(),
		reqproto.GetActionId())
	resp, err := p.privilegeEndpoints.PrivilegeActionCreate(ctx, privilegeactionCreateReqProto)
	if err != nil {
		errMsg := tlog.Error("privilegeaction new (%v) err (privilege svc rpc %v).", reqproto, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeactionNewRespProto := &privilegeproto.PrivilegeActionNewRespProto{
			ErrCode: errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeactionNewRespProto, nil
	}

	privilegeactionNewRespProto := &privilegeproto.PrivilegeActionNewRespProto{
		ErrCode:         errconst.COMMON_API_ERROR_OK,
		PrivilegeAction: resp.GetPrivilegeAction(),
	}

	return privilegeactionNewRespProto, nil
}
