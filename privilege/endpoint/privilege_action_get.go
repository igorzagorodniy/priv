package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeActionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionGetReqProto)

	return reqProto, nil
}

func DecodePrivilegeActionGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionGetReqProto)

	return reqProto, nil
}

func EncodePrivilegeActionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionGetRespProto)

	return respProto, nil
}

func DecodePrivilegeActionGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionGetRespProto)

	return respProto, nil
}

func MakePrivilegeActionGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeActionGetReqProto)

		respProto, err := service.PrivilegeActionGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeActionGet(ctx context.Context, reqproto *privilegeproto.PrivilegeActionGetReqProto) (*privilegeproto.PrivilegeActionGetRespProto, error) {
	resp, err := u.PrivilegeActionGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeActionGetRespProto), nil
}
