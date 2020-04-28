package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RolePrivilegeDel(ctx context.Context, req *privilegeproto.RolePrivilegeDelReqProto) (*privilegeproto.RolePrivilegeDelRespProto, error) {
	_, resp, err := p.rolePrivilegeDelHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RolePrivilegeDelRespProto)

	return respProto, nil
}
