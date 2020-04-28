package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeActionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionGetReqProto)

	return reqProto, nil
}

func DecodeActionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionGetReqProto)

	return reqProto, nil
}

func EncodeActionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionGetRespProto)

	return respProto, nil
}

func DecodeActionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionGetRespProto)

	return respProto, nil
}

func MakeActionGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.ActionGetReqProto)

		respProto, err := service.ActionGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) ActionGet(ctx context.Context, reqproto *privilegeproto.ActionGetReqProto) (*privilegeproto.ActionGetRespProto, error) {
	resp, err := u.ActionGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.ActionGetRespProto), nil
}
