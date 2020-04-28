package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeMenuGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuGetReqProto)

	return reqProto, nil
}

func DecodePrivilegeMenuGetRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuGetReqProto)

	return reqProto, nil
}

func EncodePrivilegeMenuGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuGetRespProto)

	return respProto, nil
}

func DecodePrivilegeMenuGetResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuGetRespProto)

	return respProto, nil
}

func MakePrivilegeMenuGetEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeMenuGetReqProto)

		respProto, err := service.PrivilegeMenuGet(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeMenuGet(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuGetReqProto) (*privilegeproto.PrivilegeMenuGetRespProto, error) {
	resp, err := u.PrivilegeMenuGetEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeMenuGetRespProto), nil
}
