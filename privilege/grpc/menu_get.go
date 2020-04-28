package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) MenuGet(ctx context.Context, req *privilegeproto.MenuGetReqProto) (*privilegeproto.MenuGetRespProto, error) {
	_, resp, err := p.menuGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.MenuGetRespProto)

	return respProto, nil
}
