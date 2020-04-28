package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) ActionNew(ctx context.Context, req *privilegeproto.ActionNewReqProto) (*privilegeproto.ActionNewRespProto, error) {
	_, resp, err := p.actionNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.ActionNewRespProto)

	return respProto, nil
}
