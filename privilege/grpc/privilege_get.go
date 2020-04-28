package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeGet(ctx context.Context, req *privilegeproto.PrivilegeGetReqProto) (*privilegeproto.PrivilegeGetRespProto, error) {
	_, resp, err := p.privilegeGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeGetRespProto)

	return respProto, nil
}
