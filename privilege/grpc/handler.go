package grpc

import (
	"xtech-kit/api/privilege-api/pkg/privilege/endpoint"
	privilegeproto "xtech-kit/proto/privilege/v1"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	actionNewHandler  kitgrpc.Handler
	actionGetHandler  kitgrpc.Handler
	actionEditHandler kitgrpc.Handler

	privilegeNewHandler  kitgrpc.Handler
	privilegeGetHandler  kitgrpc.Handler
	privilegeEditHandler kitgrpc.Handler

	roleNewHandler  kitgrpc.Handler
	roleGetHandler  kitgrpc.Handler
	roleEditHandler kitgrpc.Handler

	menuNewHandler  kitgrpc.Handler
	menuGetHandler  kitgrpc.Handler
	menuEditHandler kitgrpc.Handler

	privilegeActionNewHandler kitgrpc.Handler
	privilegeActionGetHandler kitgrpc.Handler
	privilegeActionDelHandler kitgrpc.Handler

	rolePrivilegeNewHandler kitgrpc.Handler
	rolePrivilegeGetHandler kitgrpc.Handler
	rolePrivilegeDelHandler kitgrpc.Handler

	privilegeMenuNewHandler kitgrpc.Handler
	privilegeMenuGetHandler kitgrpc.Handler
	privilegeMenuDelHandler kitgrpc.Handler

	permissionGetHandler kitgrpc.Handler
}

func NewGRPCServer(endpoints endpoint.PrivilegeApiEndpoints) privilegeproto.PrivilegeApiServiceServer {
	options := []kitgrpc.ServerOption{
		//kitgrpc.ServerErrorLogger()
	}

	// action
	actionNewHandler := kitgrpc.NewServer(
		endpoints.ActionNewEndpoint,
		endpoint.DecodeActionNewRequest,
		endpoint.EncodeActionNewResponse,
		options...,
	)
	actionGetHandler := kitgrpc.NewServer(
		endpoints.ActionGetEndpoint,
		endpoint.DecodeActionGetRequest,
		endpoint.EncodeActionGetResponse,
		options...,
	)
	actionEditHandler := kitgrpc.NewServer(
		endpoints.ActionEditEndpoint,
		endpoint.DecodeActionEditRequest,
		endpoint.EncodeActionEditResponse,
		options...,
	)

	// privilege
	privilegeNewHandler := kitgrpc.NewServer(
		endpoints.PrivilegeNewEndpoint,
		endpoint.DecodePrivilegeNewRequest,
		endpoint.EncodePrivilegeNewResponse,
		options...,
	)
	privilegeGetHandler := kitgrpc.NewServer(
		endpoints.PrivilegeGetEndpoint,
		endpoint.DecodePrivilegeGetRequest,
		endpoint.EncodePrivilegeGetResponse,
		options...,
	)
	privilegeEditHandler := kitgrpc.NewServer(
		endpoints.PrivilegeEditEndpoint,
		endpoint.DecodePrivilegeEditRequest,
		endpoint.EncodePrivilegeEditResponse,
		options...,
	)

	menuNewHandler := kitgrpc.NewServer(
		endpoints.MenuNewEndpoint,
		endpoint.DecodeMenuNewRequest,
		endpoint.EncodeMenuNewResponse,
		options...,
	)
	menuGetHandler := kitgrpc.NewServer(
		endpoints.MenuGetEndpoint,
		endpoint.DecodeMenuGetRequest,
		endpoint.EncodeMenuGetResponse,
		options...,
	)
	menuEditHandler := kitgrpc.NewServer(
		endpoints.MenuEditEndpoint,
		endpoint.DecodeMenuEditRequest,
		endpoint.EncodeMenuEditResponse,
		options...,
	)

	// role
	roleNewHandler := kitgrpc.NewServer(
		endpoints.RoleNewEndpoint,
		endpoint.DecodeRoleNewRequest,
		endpoint.EncodeRoleNewResponse,
		options...,
	)
	roleGetHandler := kitgrpc.NewServer(
		endpoints.RoleGetEndpoint,
		endpoint.DecodeRoleGetRequest,
		endpoint.EncodeRoleGetResponse,
		options...,
	)
	roleEditHandler := kitgrpc.NewServer(
		endpoints.RoleEditEndpoint,
		endpoint.DecodeRoleEditRequest,
		endpoint.EncodeRoleEditResponse,
		options...,
	)

	// privilege action
	privilegeActionNewHandler := kitgrpc.NewServer(
		endpoints.PrivilegeActionNewEndpoint,
		endpoint.DecodePrivilegeActionNewRequest,
		endpoint.EncodePrivilegeActionNewResponse,
		options...,
	)
	privilegeActionGetHandler := kitgrpc.NewServer(
		endpoints.PrivilegeActionGetEndpoint,
		endpoint.DecodePrivilegeActionGetRequest,
		endpoint.EncodePrivilegeActionGetResponse,
		options...,
	)
	privilegeActionDelHandler := kitgrpc.NewServer(
		endpoints.PrivilegeActionDelEndpoint,
		endpoint.DecodePrivilegeActionDelRequest,
		endpoint.EncodePrivilegeActionDelResponse,
		options...,
	)

	// privilege menu
	privilegeMenuNewHandler := kitgrpc.NewServer(
		endpoints.PrivilegeMenuNewEndpoint,
		endpoint.DecodePrivilegeMenuNewRequest,
		endpoint.EncodePrivilegeMenuNewResponse,
		options...,
	)
	privilegeMenuGetHandler := kitgrpc.NewServer(
		endpoints.PrivilegeMenuGetEndpoint,
		endpoint.DecodePrivilegeMenuGetRequest,
		endpoint.EncodePrivilegeMenuGetResponse,
		options...,
	)
	privilegeMenuDelHandler := kitgrpc.NewServer(
		endpoints.PrivilegeMenuDelEndpoint,
		endpoint.DecodePrivilegeMenuDelRequest,
		endpoint.EncodePrivilegeMenuDelResponse,
		options...,
	)

	// role privilege
	rolePrivilegeNewHandler := kitgrpc.NewServer(
		endpoints.RolePrivilegeNewEndpoint,
		endpoint.DecodeRolePrivilegeNewRequest,
		endpoint.EncodeRolePrivilegeNewResponse,
		options...,
	)
	rolePrivilegeGetHandler := kitgrpc.NewServer(
		endpoints.RolePrivilegeGetEndpoint,
		endpoint.DecodeRolePrivilegeGetRequest,
		endpoint.EncodeRolePrivilegeGetResponse,
		options...,
	)
	rolePrivilegeDelHandler := kitgrpc.NewServer(
		endpoints.RolePrivilegeDelEndpoint,
		endpoint.DecodeRolePrivilegeDelRequest,
		endpoint.EncodeRolePrivilegeDelResponse,
		options...,
	)

	permissionGetHandler := kitgrpc.NewServer(
		endpoints.PermissionGetEndpoint,
		endpoint.DecodePermissionGetRequest,
		endpoint.EncodePermissionGetResponse,
		options...,
	)

	var server privilegeproto.PrivilegeApiServiceServer

	server = &grpcServer{
		privilegeNewHandler:  privilegeNewHandler,
		privilegeGetHandler:  privilegeGetHandler,
		privilegeEditHandler: privilegeEditHandler,

		actionNewHandler:  actionNewHandler,
		actionGetHandler:  actionGetHandler,
		actionEditHandler: actionEditHandler,

		menuNewHandler:  menuNewHandler,
		menuGetHandler:  menuGetHandler,
		menuEditHandler: menuEditHandler,

		roleNewHandler:  roleNewHandler,
		roleGetHandler:  roleGetHandler,
		roleEditHandler: roleEditHandler,

		privilegeActionNewHandler: privilegeActionNewHandler,
		privilegeActionGetHandler: privilegeActionGetHandler,
		privilegeActionDelHandler: privilegeActionDelHandler,

		privilegeMenuNewHandler: privilegeMenuNewHandler,
		privilegeMenuGetHandler: privilegeMenuGetHandler,
		privilegeMenuDelHandler: privilegeMenuDelHandler,

		rolePrivilegeNewHandler: rolePrivilegeNewHandler,
		rolePrivilegeGetHandler: rolePrivilegeGetHandler,
		rolePrivilegeDelHandler: rolePrivilegeDelHandler,

		permissionGetHandler: permissionGetHandler,
	}

	return server
}
