package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) MenuEdit(ctx context.Context, req *privilegeproto.MenuEditReqProto) (*privilegeproto.MenuEditRespProto, error) {
	_, resp, err := p.menuEditHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.MenuEditRespProto)

	return respProto, nil
}
