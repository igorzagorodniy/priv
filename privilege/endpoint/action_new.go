package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeActionNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionNewReqProto)

	return reqProto, nil
}

func DecodeActionNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionNewReqProto)

	return reqProto, nil
}

func EncodeActionNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionNewRespProto)

	return respProto, nil
}

func DecodeActionNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionNewRespProto)

	return respProto, nil
}

func MakeActionNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.ActionNewReqProto)

		respProto, err := service.ActionNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) ActionNew(ctx context.Context, reqproto *privilegeproto.ActionNewReqProto) (*privilegeproto.ActionNewRespProto, error) {
	resp, err := u.ActionNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.ActionNewRespProto), nil
}
