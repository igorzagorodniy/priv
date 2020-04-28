package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) ActionGet(ctx context.Context, req *privilegeproto.ActionGetReqProto) (*privilegeproto.ActionGetRespProto, error) {
	_, resp, err := p.actionGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.ActionGetRespProto)

	return respProto, nil
}
