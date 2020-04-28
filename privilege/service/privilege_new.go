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

func (p *basicPrivilegeApiService) PrivilegeNew(ctx context.Context, reqproto *privilegeproto.PrivilegeNewReqProto) (*privilegeproto.PrivilegeNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilege new (%s, %s) err (recover %v).",
				reqproto.GetParentId(), reqproto.GetName(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilege new"), errMsg))
		}
	}()

	privilegeCreateReqProto := privilegecligrpc.CreatePrivilegeCreateReqProto(reqproto.ParentId, reqproto.Name)
	resp, err := p.privilegeEndpoints.PrivilegeCreate(ctx, privilegeCreateReqProto)
	if err != nil {
		errMsg := tlog.Error("privilege new (%d, %s) err (privilege svc rpc %v).", reqproto.ParentId, reqproto.Name, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeNewRespProto := &privilegeproto.PrivilegeNewRespProto{
			ErrCode: errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeNewRespProto, nil
	}

	privilegeNewRespProto := &privilegeproto.PrivilegeNewRespProto{
		ErrCode:   errconst.COMMON_API_ERROR_OK,
		Privilege: resp.Privilege,
	}

	return privilegeNewRespProto, nil
}
