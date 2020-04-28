package service

import (
	"context"
	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
	"time"
	"xtech-kit/const"
	"xtech-kit/error"
	privilegeproto "xtech-kit/proto/privilege/v1"
	privilegecligrpc "xtech-kit/svc/privilege-svc/client/grpc"
)

func (p *basicPrivilegeApiService) PrivilegeActionGet(ctx context.Context, reqproto *privilegeproto.PrivilegeActionGetReqProto) (*privilegeproto.PrivilegeActionGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilegeaction get (%v) err (recover %v).",
				reqproto, logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilegeaction get"), errMsg))
		}
	}()

	// TODO:业务逻辑判断

	privilegeactionFetchReqProto := privilegecligrpc.CreatePrivilegeActionFetchReqProto(

		0,
		reqproto.GetPrivilegeId(),
		0,
		time.Now().Unix(),
		time.Now().Unix(), reqproto.From, reqproto.Limit, reqproto.OrderBy, privilegeproto.SORT_SORT_ASC, reqproto.FieldRange)
	resp, err := p.privilegeEndpoints.PrivilegeActionFetch(ctx, privilegeactionFetchReqProto)
	if err != nil {
		errMsg := tlog.Error("privilegeaction get (%v) err (privilegeaction svc rpc %v).", reqproto, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeactionGetRespProto := &privilegeproto.PrivilegeActionGetRespProto{
			ErrCode: errconst.PARTNER_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeactionGetRespProto, nil
	}

	privilegeactionGetRespProto := &privilegeproto.PrivilegeActionGetRespProto{
		ErrCode:         errconst.COMMON_API_ERROR_OK,
		PrivilegeAction: resp.GetPrivilegeAction(),
		Total:           resp.GetTotal(),
	}

	return privilegeactionGetRespProto, nil
}
