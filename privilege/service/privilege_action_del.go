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

func (p *basicPrivilegeApiService) PrivilegeActionDel(ctx context.Context, reqproto *privilegeproto.PrivilegeActionDelReqProto) (*privilegeproto.PrivilegeActionDelRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("privilegeaction del (%d) err (recover %v).",
				reqproto.GetId(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("privilegeaction del"), errMsg))
		}
	}()

	// TODO:业务逻辑判断

	// 1. check exist
	fieldRange := make(map[string]bool, 0)
	fieldRange["id"] = true
	privilegeactionFetchReqProto := privilegecligrpc.CreatePrivilegeActionFetchReqProto(

		reqproto.GetId(),
		0,
		0,
		time.Now().Unix(),
		time.Now().Unix(), 0, 1, "", privilegeproto.SORT_SORT_ASC, fieldRange)
	exist, err := p.privilegeEndpoints.PrivilegeActionFetch(ctx, privilegeactionFetchReqProto)
	if err != nil {
		errMsg := tlog.Error("privilegeaction del (%v) err (privilege svc rpc %v).", reqproto, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
			ErrCode: errconst.PARTNER_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeactionDelRespProto, nil
	}
	if exist.ErrCode == errconst.COMMON_API_ERROR_OK {
		if len(exist.GetPrivilegeAction()) <= 0 {
			privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
				ErrCode: errconst.COMMON_API_ERROR_OBJECT_NOT_EXIST,
				ErrMsg:  "PrivilegeAction not exist",
			}
			return privilegeactionDelRespProto, nil
		}
	} else {
		privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
			ErrCode: exist.ErrCode,
			ErrMsg:  exist.ErrMsg,
		}
		return privilegeactionDelRespProto, nil
	}

	// 2. remove
	privilegeactionRemoveReqProto := privilegecligrpc.CreatePrivilegeActionRemoveReqProto(reqproto.GetId())
	resp, err := p.privilegeEndpoints.PrivilegeActionRemove(ctx, privilegeactionRemoveReqProto)
	if err != nil {
		errMsg := tlog.Error("privilegeaction del (%v) err (privilege svc rpc %v).", reqproto, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
			ErrCode: errconst.PARTNER_API_ERROR_SERVER_ABNORMAL,
			ErrMsg:  errMsg,
		}

		return privilegeactionDelRespProto, nil
	}

	if resp.ErrCode == errconst.COMMON_API_ERROR_OK {
		privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
			ErrCode: errconst.COMMON_API_ERROR_OK,
			ErrMsg:  "",
		}
		return privilegeactionDelRespProto, nil
	} else {
		privilegeactionDelRespProto := &privilegeproto.PrivilegeActionDelRespProto{
			ErrCode: resp.ErrCode,
			ErrMsg:  resp.ErrMsg,
		}
		return privilegeactionDelRespProto, nil
	}

}
