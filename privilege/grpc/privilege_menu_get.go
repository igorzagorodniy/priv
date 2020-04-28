package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PrivilegeMenuGet(ctx context.Context, req *privilegeproto.PrivilegeMenuGetReqProto) (*privilegeproto.PrivilegeMenuGetRespProto, error) {
	_, resp, err := p.privilegeMenuGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PrivilegeMenuGetRespProto)

	return respProto, nil
}
