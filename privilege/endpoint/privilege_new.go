package endpoint

import (
	"context"
	"log"
	"xtech-kit/api/privilege-api/pkg/privilege/service"
	"xtech-kit/proto/privilege/v1"

	"github.com/go-kit/kit/endpoint"
)

func EncodePrivilegeNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeNewReqProto)

	return reqProto, nil
}

func DecodePrivilegeNewRequest(ctx context.Context, req interface{}) (interface{}, error) {
	reqProto := req.(*privilegeproto.PrivilegeNewReqProto)

	return reqProto, nil
}

func EncodePrivilegeNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeNewRespProto)

	return respProto, nil
}

func DecodePrivilegeNewResponse(ctx context.Context, resp interface{}) (interface{}, error) {
	respProto := resp.(*privilegeproto.PrivilegeNewRespProto)

	return respProto, nil
}

func MakePrivilegeNewEndpoint(service service.PrivilegeApiService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		reqProto := req.(*privilegeproto.PrivilegeNewReqProto)

		log.Printf("role new 1")
		respProto, err := service.PrivilegeNew(ctx, reqProto)

		if err != nil {
			return nil, err
		}

		return respProto, nil
	}
}

// for grpc client
func (u PrivilegeApiEndpoints) PrivilegeNew(ctx context.Context, reqproto *privilegeproto.PrivilegeNewReqProto) (*privilegeproto.PrivilegeNewRespProto, error) {
	log.Printf("role new 2")
	resp, err := u.PrivilegeNewEndpoint(ctx, reqproto)

	if err != nil {
		return nil, err
	}

	return resp.(*privilegeproto.PrivilegeNewRespProto), nil
}
