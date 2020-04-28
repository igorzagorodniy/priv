package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RolePrivilegeNew(ctx context.Context, req *privilegeproto.RolePrivilegeNewReqProto) (*privilegeproto.RolePrivilegeNewRespProto, error) {
	_, resp, err := p.rolePrivilegeNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RolePrivilegeNewRespProto)

	return respProto, nil
}
