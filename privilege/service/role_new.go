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

func (p *basicPrivilegeApiService) RoleNew(ctx context.Context, reqproto *privilegeproto.RoleNewReqProto) (*privilegeproto.RoleNewRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("role new (%s err (recover %v).",
				reqproto.GetName(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("role new"), errMsg))
		}
	}()

	roleCreateReqProto := privilegecligrpc.CreateRoleCreateReqProto(reqproto.ParentId, reqproto.Name)
	resp, err := p.privilegeEndpoints.RoleCreate(ctx, roleCreateReqProto)
	if err != nil {
		errMsg := tlog.Error("role new (%d, %s) err (privilege svc rpc %v).", reqproto.ParentId, reqproto.Name, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		roleNewRespProto := &privilegeproto.RoleNewRespProto{
			ErrCode: errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return roleNewRespProto, nil
	}

	roleNewRespProto := &privilegeproto.RoleNewRespProto{
		ErrCode: errconst.COMMON_API_ERROR_OK,
		Role:    resp.GetRole(),
	}

	return roleNewRespProto, nil
}
