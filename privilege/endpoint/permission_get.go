package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePermissionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PermissionGetReqProto)

	return reqProto, nil
}

func DecodePermissionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PermissionGetReqProto)

	return reqProto, nil
}

func EncodePermissionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PermissionGetRespProto)

	return respProto, nil
}

func DecodePermissionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PermissionGetRespProto)

	return respProto, nil
}

func MakePermissionGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PermissionGetReqProto)

		respProto, err := service.PermissionGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PermissionGet(ctx context.Context, reqproto *privilegeproto.PermissionGetReqProto) (*privilegeproto.PermissionGetRespProto, error) {
	resp, err := u.PermissionGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PermissionGetRespProto), nil
}
