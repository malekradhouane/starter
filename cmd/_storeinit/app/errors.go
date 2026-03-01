package app

import "errors"

// sentinel error: do not use alone.
var ErrApplication = errors.New("app")
