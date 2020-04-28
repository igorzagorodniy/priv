package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeRoleNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleNewReqProto)

	return reqProto, nil
}

func DecodeRoleNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.RoleNewReqProto)

	return reqProto, nil
}

func EncodeRoleNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleNewRespProto)

	return respProto, nil
}

func DecodeRoleNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.RoleNewRespProto)

	return respProto, nil
}

func MakeRoleNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.RoleNewReqProto)

		respProto, err := service.RoleNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) RoleNew(ctx context.Context, reqproto *privilegeproto.RoleNewReqProto) (*privilegeproto.RoleNewRespProto, error) {
	resp, err := u.RoleNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.RoleNewRespProto), nil
}
