package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeActionGet(ctx context.Context, req *privilegeproto.PrivilegeActionGetReqProto) (*privilegeproto.PrivilegeActionGetRespProto, error) {
	_, resp, err := p.privilegeActionGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeActionGetRespProto)

	return respProto, nil
}
