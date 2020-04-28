package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeActionNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionNewReqProto)

	return reqProto, nil
}

func DecodePrivilegeActionNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionNewReqProto)

	return reqProto, nil
}

func EncodePrivilegeActionNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionNewRespProto)

	return respProto, nil
}

func DecodePrivilegeActionNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionNewRespProto)

	return respProto, nil
}

func MakePrivilegeActionNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeActionNewReqProto)

		respProto, err := service.PrivilegeActionNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeActionNew(ctx context.Context, reqproto *privilegeproto.PrivilegeActionNewReqProto) (*privilegeproto.PrivilegeActionNewRespProto, error) {
	resp, err := u.PrivilegeActionNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeActionNewRespProto), nil
}
