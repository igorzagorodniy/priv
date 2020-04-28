package endpoint

import (
	"xtech-kit/api/privilege-api/pkg/privilege/service"

	kitendpoint "github.com/go-kit/kit/endpoint"
)

type PrivilegeApiEndpoints struct {
	ActionNewEndpoint  kitendpoint.Endpoint
	ActionGetEndpoint  kitendpoint.Endpoint
	ActionEditEndpoint kitendpoint.Endpoint

	PrivilegeNewEndpoint  kitendpoint.Endpoint
	PrivilegeGetEndpoint  kitendpoint.Endpoint
	PrivilegeEditEndpoint kitendpoint.Endpoint

	RoleNewEndpoint  kitendpoint.Endpoint
	RoleGetEndpoint  kitendpoint.Endpoint
	RoleEditEndpoint kitendpoint.Endpoint

	MenuNewEndpoint  kitendpoint.Endpoint
	MenuGetEndpoint  kitendpoint.Endpoint
	MenuEditEndpoint kitendpoint.Endpoint

	PrivilegeActionNewEndpoint kitendpoint.Endpoint
	PrivilegeActionGetEndpoint kitendpoint.Endpoint
	PrivilegeActionDelEndpoint kitendpoint.Endpoint

	PrivilegeMenuNewEndpoint kitendpoint.Endpoint
	PrivilegeMenuGetEndpoint kitendpoint.Endpoint
	PrivilegeMenuDelEndpoint kitendpoint.Endpoint

	RolePrivilegeNewEndpoint kitendpoint.Endpoint
	RolePrivilegeGetEndpoint kitendpoint.Endpoint
	RolePrivilegeDelEndpoint kitendpoint.Endpoint

	PermissionGetEndpoint kitendpoint.Endpoint
}

func NewPrivilegeApiEndpoints(service service.PrivilegeApiService) PrivilegeApiEndpoints {
	return PrivilegeApiEndpoints{
		PrivilegeNewEndpoint:  MakePrivilegeNewEndpoint(service),
		PrivilegeGetEndpoint:  MakePrivilegeGetEndpoint(service),
		PrivilegeEditEndpoint: MakePrivilegeEditEndpoint(service),

		ActionNewEndpoint:  MakeActionNewEndpoint(service),
		ActionGetEndpoint:  MakeActionGetEndpoint(service),
		ActionEditEndpoint: MakeActionEditEndpoint(service),

		MenuNewEndpoint:  MakeMenuNewEndpoint(service),
		MenuGetEndpoint:  MakeMenuGetEndpoint(service),
		MenuEditEndpoint: MakeMenuEditEndpoint(service),

		RoleNewEndpoint:  MakeRoleNewEndpoint(service),
		RoleGetEndpoint:  MakeRoleGetEndpoint(service),
		RoleEditEndpoint: MakeRoleEditEndpoint(service),

		PrivilegeActionNewEndpoint: MakePrivilegeActionNewEndpoint(service),
		PrivilegeActionGetEndpoint: MakePrivilegeActionGetEndpoint(service),
		PrivilegeActionDelEndpoint: MakePrivilegeActionDelEndpoint(service),

		PrivilegeMenuNewEndpoint: MakePrivilegeMenuNewEndpoint(service),
		PrivilegeMenuGetEndpoint: MakePrivilegeMenuGetEndpoint(service),
		PrivilegeMenuDelEndpoint: MakePrivilegeMenuDelEndpoint(service),

		RolePrivilegeNewEndpoint: MakeRolePrivilegeNewEndpoint(service),
		RolePrivilegeGetEndpoint: MakeRolePrivilegeGetEndpoint(service),
		RolePrivilegeDelEndpoint: MakeRolePrivilegeDelEndpoint(service),

		PermissionGetEndpoint: MakePermissionGetEndpoint(service),
	}
}
