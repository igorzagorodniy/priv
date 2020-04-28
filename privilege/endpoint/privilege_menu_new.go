package endpoint

import (
	"context"

	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeMenuNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuNewReqProto)

	return reqProto, nil
}

func DecodePrivilegeMenuNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeMenuNewReqProto)

	return reqProto, nil
}

func EncodePrivilegeMenuNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuNewRespProto)

	return respProto, nil
}

func DecodePrivilegeMenuNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeMenuNewRespProto)

	return respProto, nil
}

func MakePrivilegeMenuNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeMenuNewReqProto)

		respProto, err := service.PrivilegeMenuNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeMenuNew(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuNewReqProto) (*privilegeproto.PrivilegeMenuNewRespProto, error) {
	resp, err := u.PrivilegeMenuNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeMenuNewRespProto), nil
}
