package core

import "log/slog"

type LoggerContract interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	With(...any) *slog.Logger
	WithGroup(string) *slog.Logger
}

var _ LoggerContract = (*slog.Logger)(nil)
