package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RoleNew(ctx context.Context, req *privilegeproto.RoleNewReqProto) (*privilegeproto.RoleNewRespProto, error) {
	_, resp, err := p.roleNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RoleNewRespProto)

	return respProto, nil
}
