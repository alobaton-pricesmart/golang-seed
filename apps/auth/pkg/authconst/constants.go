package authconst

const (
	GeneralErrorInternalServerError    = "general.error.internalServerError"
	GeneralErrorMarshal                = "general.error.marshal"
	GeneralErrorRegisterNotFound       = "general.error.registerNotFound"
	GeneralErrorRegisterNotFoundParams = "general.error.registerNotFoundParams"
	GeneralErrorRequiredField          = "general.error.requiredField"
	GeneralErrorInvalidField           = "general.error.invalidField"
	GeneralErrorAccessingDatabase      = "general.error.accessingDatabase"
	GeneralErrorRegisterAlreadyExists  = "general.error.registerAlreadyExists"

	OAuthInvalidUsernamePassword = "oauth.invalidUsernamePassword"
	OauthInvalidToken            = "oauth.invalidToken"
	OAuthPermissionDenied        = "oauth.permissionDenied"

	UsersUser  = "users.user"
	UsersUsers = "users.users"

	ClientsClient  = "clients.client"
	ClientsClients = "clients.clients"

	PermissionsPermission  = "permissions.permission"
	PermissionsPermissions = "permissions.permissions"

	RolesRole  = "roles.role"
	RolesRoles = "roles.roles"
)
