package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RolePrivilegeGet(ctx context.Context, req *privilegeproto.RolePrivilegeGetReqProto) (*privilegeproto.RolePrivilegeGetRespProto, error) {
	_, resp, err := p.rolePrivilegeGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RolePrivilegeGetRespProto)

	return respProto, nil
}
