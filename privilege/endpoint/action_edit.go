package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodeActionEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionEditReqProto)

	return reqProto, nil
}

func DecodeActionEditRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.ActionEditReqProto)

	return reqProto, nil
}

func EncodeActionEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionEditRespProto)

	return respProto, nil
}

func DecodeActionEditResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.ActionEditRespProto)

	return respProto, nil
}

func MakeActionEditEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.ActionEditReqProto)

		respProto, err := service.ActionEdit(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) ActionEdit(ctx context.Context, reqproto *privilegeproto.ActionEditReqProto) (*privilegeproto.ActionEditRespProto, error) {
	resp, err := u.ActionEditEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.ActionEditRespProto), nil
}
