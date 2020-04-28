package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeActionNew(ctx context.Context, req *privilegeproto.PrivilegeActionNewReqProto) (*privilegeproto.PrivilegeActionNewRespProto, error) {
	_, resp, err := p.privilegeActionNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeActionNewRespProto)

	return respProto, nil
}
