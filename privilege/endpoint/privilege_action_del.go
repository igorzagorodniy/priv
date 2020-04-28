package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeActionDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionDelReqProto)

	return reqProto, nil
}

func DecodePrivilegeActionDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeActionDelReqProto)

	return reqProto, nil
}

func EncodePrivilegeActionDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionDelRespProto)

	return respProto, nil
}

func DecodePrivilegeActionDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeActionDelRespProto)

	return respProto, nil
}

func MakePrivilegeActionDelEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeActionDelReqProto)

		respProto, err := service.PrivilegeActionDel(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeActionDel(ctx context.Context, reqproto *privilegeproto.PrivilegeActionDelReqProto) (*privilegeproto.PrivilegeActionDelRespProto, error) {
	resp, err := u.PrivilegeActionDelEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeActionDelRespProto), nil
}
