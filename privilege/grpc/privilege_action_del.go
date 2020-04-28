package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeActionDel(ctx context.Context, req *privilegeproto.PrivilegeActionDelReqProto) (*privilegeproto.PrivilegeActionDelRespProto, error) {
	_, resp, err := p.privilegeActionDelHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeActionDelRespProto)

	return respProto, nil
}
