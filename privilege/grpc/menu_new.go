package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) MenuNew(ctx context.Context, req *privilegeproto.MenuNewReqProto) (*privilegeproto.MenuNewRespProto, error) {
	_, resp, err := p.menuNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.MenuNewRespProto)

	return respProto, nil
}
