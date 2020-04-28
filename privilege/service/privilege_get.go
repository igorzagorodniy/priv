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

func (p *basicPrivilegeApiService) PrivilegeGet(ctx context.Context, reqproto *privilegeproto.PrivilegeGetReqProto) (*privilegeproto.PrivilegeGetRespProto, error) {

	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilege get (%s, %d) err (recover %v).",
				reqproto.GetName(), reqproto.GetParentId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilege get"), errMsg))
		}
	}()

	// todo : 校验权限

	privilegeFetchReqProto := privilegecligrpc.CreatePrivilegeFetchReqProto(reqproto.ParentId, reqproto.Name, reqproto.From, reqproto.Limit)
	resp, err := p.privilegeEndpoints.PrivilegeFetch(ctx, privilegeFetchReqProto)
	if err != nil {
		errMsg := tlog.Error("privilege fetch (%d, %s) err (privilege svc rpc %v).", reqproto.ParentId, reqproto.Name, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeGetRespProto := &privilegeproto.PrivilegeGetRespProto{
			ErrCode: errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeGetRespProto, nil
	}
	privilegeGetRespProto := &privilegeproto.PrivilegeGetRespProto{
		ErrCode:   errconst.COMMON_API_ERROR_OK,
		Privilege: resp.Privilege,
	}

	return privilegeGetRespProto, nil
}
