package app

import (
	"fmt"

	"storeinit/domain"
)

func NewApplicationBuilder() *builder {
	return &builder{}
}

type builder struct {
	Application
}

func (app *builder) Build() (*Application, error) {
	if len(app.repos) == 0 {
		return nil, fmt.Errorf("%w: repository is missing", ErrApplication)
	}

	return &app.Application, nil
}

func (app *builder) SetRepositories(repos ...domain.RepositoryContract) *builder {
	app.repos = repos

	return app
}
