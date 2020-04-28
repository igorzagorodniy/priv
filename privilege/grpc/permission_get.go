package grpc

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
)

func (p *grpcServer) PermissionGet(ctx context.Context, req *privilegeproto.PermissionGetReqProto) (*privilegeproto.PermissionGetRespProto, error) {
	_, resp, err := p.permissionGetHandler.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	respProto := resp.(*privilegeproto.PermissionGetRespProto)

	return respProto, nil
}
