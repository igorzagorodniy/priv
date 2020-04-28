package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeMenuNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuNewReqProto)

	return reqProto, nil
}

func DecodeMenuNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.MenuNewReqProto)

	return reqProto, nil
}

func EncodeMenuNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuNewRespProto)

	return respProto, nil
}

func DecodeMenuNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.MenuNewRespProto)

	return respProto, nil
}

func MakeMenuNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.MenuNewReqProto)

		respProto, err := service.MenuNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) MenuNew(ctx context.Context, reqproto *privilegeproto.MenuNewReqProto) (*privilegeproto.MenuNewRespProto, error) {
	resp, err := u.MenuNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.MenuNewRespProto), nil
}
