package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRolePrivilegeDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeDelReqProto)

	return reqProto, nil
}

func DecodeRolePrivilegeDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RolePrivilegeDelReqProto)

	return reqProto, nil
}

func EncodeRolePrivilegeDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeDelRespProto)

	return respProto, nil
}

func DecodeRolePrivilegeDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RolePrivilegeDelRespProto)

	return respProto, nil
}

func MakeRolePrivilegeDelEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RolePrivilegeDelReqProto)

		respProto, err := service.RolePrivilegeDel(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RolePrivilegeDel(ctx context.Context, reqproto *privilegeproto.RolePrivilegeDelReqProto) (*privilegeproto.RolePrivilegeDelRespProto, error) {
	resp, err := u.RolePrivilegeDelEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RolePrivilegeDelRespProto), nil
}
