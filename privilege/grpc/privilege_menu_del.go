package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeMenuDel(ctx context.Context, req *privilegeproto.PrivilegeMenuDelReqProto) (*privilegeproto.PrivilegeMenuDelRespProto, error) {
	_, resp, err := p.privilegeMenuDelHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeMenuDelRespProto)

	return respProto, nil
}
