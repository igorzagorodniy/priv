package endpoint

import (
	"context"
	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeGetReqProto)

	return reqProto, nil
}

func DecodePrivilegeGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeGetReqProto)

	return reqProto, nil
}

func EncodePrivilegeGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeGetRespProto)

	return respProto, nil
}

func DecodePrivilegeGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeGetRespProto)

	return respProto, nil
}

func MakePrivilegeGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeGetReqProto)

		respProto, err := service.PrivilegeGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeGet(ctx context.Context, reqproto *privilegeproto.PrivilegeGetReqProto) (*privilegeproto.PrivilegeGetRespProto, error) {

	resp, err := u.PrivilegeGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeGetRespProto), nil
}
