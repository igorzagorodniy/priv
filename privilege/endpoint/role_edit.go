package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRoleEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleEditReqProto)

	return reqProto, nil
}

func DecodeRoleEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleEditReqProto)

	return reqProto, nil
}

func EncodeRoleEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleEditRespProto)

	return respProto, nil
}

func DecodeRoleEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleEditRespProto)

	return respProto, nil
}

func MakeRoleEditEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RoleEditReqProto)

		respProto, err := service.RoleEdit(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RoleEdit(ctx context.Context, reqproto *privilegeproto.RoleEditReqProto) (*privilegeproto.RoleEditRespProto, error) {
	resp, err := u.RoleEditEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RoleEditRespProto), nil
}