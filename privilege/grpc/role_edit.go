package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RoleEdit(ctx context.Context, req *privilegeproto.RoleEditReqProto) (*privilegeproto.RoleEditRespProto, error) {
	_, resp, err := p.roleEditHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RoleEditRespProto)

	return respProto, nil
}
