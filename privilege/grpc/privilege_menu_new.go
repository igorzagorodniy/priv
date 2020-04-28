package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeMenuNew(ctx context.Context, req *privilegeproto.PrivilegeMenuNewReqProto) (*privilegeproto.PrivilegeMenuNewRespProto, error) {
	_, resp, err := p.privilegeMenuNewHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeMenuNewRespProto)

	return respProto, nil
}
