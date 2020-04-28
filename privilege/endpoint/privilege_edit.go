package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeEditReqProto)

	return reqProto, nil
}

func DecodePrivilegeEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeEditReqProto)

	return reqProto, nil
}

func EncodePrivilegeEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeEditRespProto)

	return respProto, nil
}

func DecodePrivilegeEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeEditRespProto)

	return respProto, nil
}

func MakePrivilegeEditEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeEditReqProto)

		respProto, err := service.PrivilegeEdit(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeEdit(ctx context.Context, reqproto *privilegeproto.PrivilegeEditReqProto) (*privilegeproto.PrivilegeEditRespProto, error) {
	resp, err := u.PrivilegeEditEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeEditRespProto), nil
}
