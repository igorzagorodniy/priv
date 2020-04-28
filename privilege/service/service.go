package service

import (
	"context"
	privilegeproto "xtech-kit/proto/privilege/v1"
	authendpoint "xtech-kit/svc/auth-svc/pkg/auth/endpoint"
	privilegeendpoint "xtech-kit/svc/privilege-svc/pkg/privilege/endpoint"
	userendpoint "xtech-kit/svc/user-svc/pkg/user/endpoint"
)

type PrivilegeApiService interface {
	ActionNew(ctx context.Context, reqproto *privilegeproto.ActionNewReqProto) (*privilegeproto.ActionNewRespProto, error)
	ActionGet(ctx context.Context, reqproto *privilegeproto.ActionGetReqProto) (*privilegeproto.ActionGetRespProto, error)
	ActionEdit(ctx context.Context, reqproto *privilegeproto.ActionEditReqProto) (*privilegeproto.ActionEditRespProto, error)

	PrivilegeNew(ctx context.Context, reqproto *privilegeproto.PrivilegeNewReqProto) (*privilegeproto.PrivilegeNewRespProto, error)
	PrivilegeGet(ctx context.Context, reqproto *privilegeproto.PrivilegeGetReqProto) (*privilegeproto.PrivilegeGetRespProto, error)
	PrivilegeEdit(ctx context.Context, reqproto *privilegeproto.PrivilegeEditReqProto) (*privilegeproto.PrivilegeEditRespProto, error)

	RoleNew(ctx context.Context, reqproto *privilegeproto.RoleNewReqProto) (*privilegeproto.RoleNewRespProto, error)
	RoleGet(ctx context.Context, reqproto *privilegeproto.RoleGetReqProto) (*privilegeproto.RoleGetRespProto, error)
	RoleEdit(ctx context.Context, reqproto *privilegeproto.RoleEditReqProto) (*privilegeproto.RoleEditRespProto, error)

	MenuNew(ctx context.Context, reqproto *privilegeproto.MenuNewReqProto) (*privilegeproto.MenuNewRespProto, error)
	MenuGet(ctx context.Context, reqproto *privilegeproto.MenuGetReqProto) (*privilegeproto.MenuGetRespProto, error)
	MenuEdit(ctx context.Context, reqproto *privilegeproto.MenuEditReqProto) (*privilegeproto.MenuEditRespProto, error)

	PrivilegeActionNew(ctx context.Context, reqproto *privilegeproto.PrivilegeActionNewReqProto) (*privilegeproto.PrivilegeActionNewRespProto, error)
	PrivilegeActionGet(ctx context.Context, reqproto *privilegeproto.PrivilegeActionGetReqProto) (*privilegeproto.PrivilegeActionGetRespProto, error)
	PrivilegeActionDel(ctx context.Context, reqproto *privilegeproto.PrivilegeActionDelReqProto) (*privilegeproto.PrivilegeActionDelRespProto, error)

	PrivilegeMenuNew(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuNewReqProto) (*privilegeproto.PrivilegeMenuNewRespProto, error)
	PrivilegeMenuGet(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuGetReqProto) (*privilegeproto.PrivilegeMenuGetRespProto, error)
	PrivilegeMenuDel(ctx context.Context, reqproto *privilegeproto.PrivilegeMenuDelReqProto) (*privilegeproto.PrivilegeMenuDelRespProto, error)

	RolePrivilegeNew(ctx context.Context, reqproto *privilegeproto.RolePrivilegeNewReqProto) (*privilegeproto.RolePrivilegeNewRespProto, error)
	RolePrivilegeGet(ctx context.Context, reqproto *privilegeproto.RolePrivilegeGetReqProto) (*privilegeproto.RolePrivilegeGetRespProto, error)
	RolePrivilegeDel(ctx context.Context, reqproto *privilegeproto.RolePrivilegeDelReqProto) (*privilegeproto.RolePrivilegeDelRespProto, error)

	PermissionGet(ctx context.Context, reqproto *privilegeproto.PermissionGetReqProto) (*privilegeproto.PermissionGetRespProto, error)
}

type basicPrivilegeApiService struct {
	userEndpoints      userendpoint.UserSvcEndpoints
	privilegeEndpoints privilegeendpoint.PrivilegeSvcEndpoints
	authEndpoints      authendpoint.AuthSvcEndpoints
}

func NewPrivilegeApiService(userendpoints userendpoint.UserSvcEndpoints, privilegeEndpoints privilegeendpoint.PrivilegeSvcEndpoints, authendpoints authendpoint.AuthSvcEndpoints) PrivilegeApiService {
	var service PrivilegeApiService

	service = &basicPrivilegeApiService{
		userEndpoints:      userendpoints,
		privilegeEndpoints: privilegeEndpoints,
		authEndpoints:      authendpoints,
	}

	// Add Middleware
	// service = LogMiddleware()(service)

	return service
}
