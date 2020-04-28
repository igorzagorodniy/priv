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

func (p *basicPrivilegeApiService) ActionNew(ctx context.Context, reqproto *privilegeproto.ActionNewReqProto) (*privilegeproto.ActionNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("action new (%s, %s) err (recover %v).",
				reqproto.GetActionName(), reqproto.GetActionPath(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("action new"), errMsg))
		}
	}()

	actionCreateReqProto := privilegecligrpc.CreateActionCreateReqProto(reqproto.GetActionName(), reqproto.GetActionPath(), reqproto.GetDescription())
	resp, err := p.privilegeEndpoints.ActionCreate(ctx, actionCreateReqProto)
	if err != nil {
		errMsg := tlog.Error("action new (%d, %s) err (privilege svc rpc %v).", reqproto.GetActionPath(), reqproto.GetActionName(), err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		actionNewRespProto := &privilegeproto.ActionNewRespProto{
			ErrCode: errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return actionNewRespProto, nil
	}

	actionNewRespProto := &privilegeproto.ActionNewRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
		Action:  resp.GetAction(),
	}

	return actionNewRespProto, nil
}
