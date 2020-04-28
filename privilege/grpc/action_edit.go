package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) ActionEdit(ctx context.Context, req *privilegeproto.ActionEditReqProto) (*privilegeproto.ActionEditRespProto, error) {
	_, resp, err := p.actionEditHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.ActionEditRespProto)

	return respProto, nil
}
