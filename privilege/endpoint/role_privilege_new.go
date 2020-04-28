package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRolePrivilegeNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeNewReqProto)

	return reqProto, nil
}

func DecodeRolePrivilegeNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeNewReqProto)

	return reqProto, nil
}

func EncodeRolePrivilegeNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeNewRespProto)

	return respProto, nil
}

func DecodeRolePrivilegeNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeNewRespProto)

	return respProto, nil
}

func MakeRolePrivilegeNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RolePrivilegeNewReqProto)

		respProto, err := service.RolePrivilegeNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RolePrivilegeNew(ctx context.Context, reqproto *privilegeproto.RolePrivilegeNewReqProto) (*privilegeproto.RolePrivilegeNewRespProto, error) {
	resp, err := u.RolePrivilegeNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RolePrivilegeNewRespProto), nil
}
