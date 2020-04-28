package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) RoleGet(ctx context.Context, req *privilegeproto.RoleGetReqProto) (*privilegeproto.RoleGetRespProto, error) {
	_, resp, err := p.roleGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.RoleGetRespProto)

	return respProto, nil
}
