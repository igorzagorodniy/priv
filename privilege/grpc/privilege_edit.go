package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeEdit(ctx context.Context, req *privilegeproto.PrivilegeEditReqProto) (*privilegeproto.PrivilegeEditRespProto, error) {
	_, resp, err := p.privilegeEditHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeEditRespProto)

	return respProto, nil
}
