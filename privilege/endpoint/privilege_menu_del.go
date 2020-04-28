package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeMenuDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuDelReqProto)

	return reqProto, nil
}

func DecodePrivilegeMenuDelRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuDelReqProto)

	return reqProto, nil
}

func EncodePrivilegeMenuDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuDelRespProto)

	return respProto, nil
}

func DecodePrivilegeMenuDelResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuDelRespProto)

	return respProto, nil
}

func MakePrivilegeMenuDelEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeMenuDelReqProto)

		respProto, err := service.PrivilegeMenuDel(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeMenuDel(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuDelReqProto) (*privilegeproto.PrivilegeMenuDelRespProto, error) {
	resp, err := u.PrivilegeMenuDelEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeMenuDelRespProto), nil
}
