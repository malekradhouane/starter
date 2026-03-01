package gatekeeper

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is intentionally minimal; wire your logger here.
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type EnforcerDependency interface {
	AddPermissionForUser(user string, permission ...string) (bool, error)
	AddRoleForUser(user string, role string, domain ...string) (bool, error)
	DeleteRoleForUser(user string, role string, domain ...string) (bool, error)
	GetPermissionsForUser(user string, domain ...string) ([][]string, error)
	DeletePermissionForUser(user string, permission ...string) (bool, error)
	Enforce(rvals ...interface{}) (bool, error)
}

type Token struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
	Expiry       time.Time
}

type IDToken interface {
	Claims(target any) error
	Subject() string
}

// IdentityStore represents how you persist the authenticated identity/tokens to the requester (cookie/session/header).
type IdentityStore interface {
	SaveIdentity(w http.ResponseWriter, r *http.Request, subject string, tok Token, claims map[string]any) error
	ClearIdentity(w http.ResponseWriter, r *http.Request) error
	SubjectFromRequest(r *http.Request) (subject string, ok bool)
	Set(ctx context.Context, token string, data SessionData, ttl time.Duration) error
	Get(ctx context.Context, token string) (*SessionData, error)
}

type CasbinHandler interface {
	Authorize(obj string, act string) gin.HandlerFunc
}

type StateStore interface {
	Save(ctx context.Context, key string, value StateData, ttl time.Duration) error
	Load(ctx context.Context, key string) (StateData, error)
	Delete(ctx context.Context, key string) error
	VerifyAndLoad(ctx context.Context, key string) (StateData, error)
	GenerateState(ctx context.Context) (string, error)
}
