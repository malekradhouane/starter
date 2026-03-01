package casbin

type LoggerContract interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

type EnforcerDependency interface {
	AddPermissionForUser(user string, permission ...string) (bool, error)
	GetPermissionsForUser(user string, domain ...string) ([][]string, error)
	DeletePermissionForUser(user string, permission ...string) (bool, error)
	Enforce(rvals ...interface{}) (bool, error)
}
