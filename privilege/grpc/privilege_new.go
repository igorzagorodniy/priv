package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeNew(ctx context.Context, req *privilegeproto.PrivilegeNewReqProto) (*privilegeproto.PrivilegeNewRespProto, error) {
	_, resp, err := p.privilegeNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeNewRespProto)

	return respProto, nil
}
