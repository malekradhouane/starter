package cli

import (
	"context"
)

type builder struct {
	Controller
}

func NewControllerBuilder() *builder {
	return &builder{}
}

func (builder *builder) Build(_ context.Context) (*Controller, error) {
	return &builder.Controller, nil
}

func (builder *builder) SetApplication(app ApplicationDeps) *builder {
	builder.app = app

	return builder
}
