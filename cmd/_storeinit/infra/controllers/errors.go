package controllers

import "errors"

// Sentinel error. Do not use alone.
var ErrController = errors.New("handler")
