package service

import (
	"context"
	"fmt"

	"xtech-kit/const"
	"xtech-kit/error"
	"xtech-kit/proto/privilege/v1"
	authcligrpc "xtech-kit/svc/auth-svc/client/grpc"

	"github.com/xianyu-tech/tlog"
	"github.com/xianyu-tech/util/log"
)

func (p *basicPrivilegeApiService) PermissionGet(ctx context.Context, reqproto *privilegeproto.PermissionGetReqProto) (*privilegeproto.PermissionGetRespProto, error) {
	defer func() {
		if rec := recover(); rec != nil {
			errMsg := tlog.Error("permission get (%s, %s) err (recover %v).",
				reqproto.GetActionName(), reqproto.GetActionPath(), logutil.PrintPanic())

			tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrPanicRecovery("permission get"), errMsg))
		}
	}()

	permissionGetRespProto := &privilegeproto.PermissionGetRespProto{}

	userId := reqproto.GetUserId()
	csrfSecret := reqproto.GetCsrfSecret()
	accessToken := reqproto.GetAccessToken()

	// 1. 校验Token合法性及有效期
	tokenFetchReqProto := authcligrpc.NewTokenFetchReqProto(userId, csrfSecret, accessToken)

	tokenFetchRespProto, err := p.authEndpoints.TokenFetch(ctx, tokenFetchReqProto)

	if err != nil {
		errMsg := tlog.Error("token fetch (%d) err (auth svc rpc %v).", userId, err)
		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))

		permissionGetRespProto.ErrCode = errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL
		permissionGetRespProto.ErrMsg = errMsg

		return permissionGetRespProto, nil
	}

	if tokenFetchRespProto.GetErrCode() != errconst.COMMON_API_ERROR_OK {
		errMsg := tlog.Error("token fetch (%d) err (token fetch %s).",
			userId, tokenFetchRespProto.GetErrMsg())
		tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrSvcExecute("token fetch"), errMsg))

		permissionGetRespProto.ErrCode = tokenFetchRespProto.GetErrCode()
		permissionGetRespProto.ErrMsg = tokenFetchRespProto.GetErrMsg()

		return permissionGetRespProto, nil
	}
	isExpired := tokenFetchRespProto.GetIsExpired()
	if isExpired == true {
		errMsg := fmt.Sprintf("token fetch (%d) err (access token expired).", userId)

		permissionGetRespProto.ErrCode = errconst.COMMON_API_ERROR_ACCESS_TOKEN_EXPIRED
		permissionGetRespProto.ErrMsg = errMsg

		return permissionGetRespProto, nil
	}

	// 2.check privilege
	actionRoleReqProto := &privilegeproto.ActionRoleFetchReqProto{ActionIndex: reqproto.ActionName + reqproto.ActionPath}
	actionRoleFetchResp, err := p.privilegeEndpoints.ActionRoleFetch(ctx, actionRoleReqProto)
	if err != nil {
		errMsg := tlog.Error("actionRole fetch (%d) err (privilege svc rpc %v).", userId, err)

		tlog.AsyncSend(tlog.NewRawLogError(err, errMsg))
		permissionGetRespProto.ErrCode = errconst.PRIVILEGE_API_ERROR_SERVER_ABNORMAL
		permissionGetRespProto.ErrMsg = errMsg

		return permissionGetRespProto, nil
	}

	if actionRoleFetchResp.GetErrCode() != errconst.COMMON_API_ERROR_OK {
		errMsg := tlog.Error("actionRole fetch (%d) err (actionRole fetch %s).",
			userId, actionRoleFetchResp.GetErrMsg())

		tlog.AsyncSend(tlog.NewRawLogError(comconst.ErrSvcExecute("actionRole fetch"), errMsg))
		permissionGetRespProto.ErrCode = actionRoleFetchResp.GetErrCode()
		permissionGetRespProto.ErrMsg = actionRoleFetchResp.GetErrMsg()

		return permissionGetRespProto, nil
	}

	tokenRoleIds := tokenFetchRespProto.GetRoldIds()
	if actionRoleFetchResp.ActionRole == nil || len(actionRoleFetchResp.ActionRole.RoleIds) <= 0 {
		permissionGetRespProto.ErrCode = errconst.SYS_GW_ERROR_PRIVILEGE_API_NO_PERMISSION
		permissionGetRespProto.Permit = false
		return permissionGetRespProto, nil
	}
	actionRoleIds := actionRoleFetchResp.ActionRole.RoleIds
	if HasIntersect(tokenRoleIds, actionRoleIds) {
		permissionGetRespProto.ErrCode = errconst.COMMON_API_ERROR_OK
		permissionGetRespProto.Permit = true
		permissionGetRespProto.RoleIds = tokenFetchRespProto.GetRoldIds()
	} else {
		permissionGetRespProto.ErrCode = errconst.SYS_GW_ERROR_PRIVILEGE_API_NO_PERMISSION
		permissionGetRespProto.Permit = false
	}

	return permissionGetRespProto, nil
}
