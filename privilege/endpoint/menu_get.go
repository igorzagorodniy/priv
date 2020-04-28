package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeMenuGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuGetReqProto)

	return reqProto, nil
}

func DecodeMenuGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuGetReqProto)

	return reqProto, nil
}

func EncodeMenuGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuGetRespProto)

	return respProto, nil
}

func DecodeMenuGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuGetRespProto)

	return respProto, nil
}

func MakeMenuGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.MenuGetReqProto)

		respProto, err := service.MenuGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) MenuGet(ctx context.Context, reqproto *privilegeproto.MenuGetReqProto) (*privilegeproto.MenuGetRespProto, error) {
	resp, err := u.MenuGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.MenuGetRespProto), nil
}
