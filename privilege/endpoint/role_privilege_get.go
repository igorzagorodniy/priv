package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRolePrivilegeGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeGetReqProto)

	return reqProto, nil
}

func DecodeRolePrivilegeGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeGetReqProto)

	return reqProto, nil
}

func EncodeRolePrivilegeGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeGetRespProto)

	return respProto, nil
}

func DecodeRolePrivilegeGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeGetRespProto)

	return respProto, nil
}

func MakeRolePrivilegeGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RolePrivilegeGetReqProto)

		respProto, err := service.RolePrivilegeGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RolePrivilegeGet(ctx context.Context, reqproto *privilegeproto.RolePrivilegeGetReqProto) (*privilegeproto.RolePrivilegeGetRespProto, error) {
	resp, err := u.RolePrivilegeGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RolePrivilegeGetRespProto), nil
}
