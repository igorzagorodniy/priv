package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRoleGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleGetReqProto)

	return reqProto, nil
}

func DecodeRoleGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleGetReqProto)

	return reqProto, nil
}

func EncodeRoleGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleGetRespProto)

	return respProto, nil
}

func DecodeRoleGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleGetRespProto)

	return respProto, nil
}

func MakeRoleGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RoleGetReqProto)

		respProto, err := service.RoleGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RoleGet(ctx context.Context, reqproto *privilegeproto.RoleGetReqProto) (*privilegeproto.RoleGetRespProto, error) {
	resp, err := u.RoleGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RoleGetRespProto), nil
}
