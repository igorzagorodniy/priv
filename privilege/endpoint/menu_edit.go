package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeMenuEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuEditReqProto)

	return reqProto, nil
}

func DecodeMenuEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuEditReqProto)

	return reqProto, nil
}

func EncodeMenuEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuEditRespProto)

	return respProto, nil
}

func DecodeMenuEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuEditRespProto)

	return respProto, nil
}

func MakeMenuEditEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.MenuEditReqProto)

		respProto, err := service.MenuEdit(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) MenuEdit(ctx context.Context, reqproto *privilegeproto.MenuEditReqProto) (*privilegeproto.MenuEditRespProto, error) {
	resp, err := u.MenuEditEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.MenuEditRespProto), nil
}
