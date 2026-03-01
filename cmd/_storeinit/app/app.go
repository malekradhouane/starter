package app

import (
	"context"
	"slices"

	"storeinit/domain"
)

type Application struct {
	repos []domain.RepositoryContract
}

func (app *Application) Close(_ context.Context) error {
	return nil
}

var _ ApplicationContract = (*Application)(nil)

func (app *Application) Migrate(ctx context.Context, req domain.MigrateRequest) []domain.MigrateResponse {
	resp := []domain.MigrateResponse{}
	repos := app.repos

	if req.Action == domain.ActionMigrateDown {
		slices.Reverse(repos)
	}

	for _, repo := range repos {
		resp = append(resp, repo.Migrate(ctx, req))
	}

	return resp
}

func (app *Application) MigrationStatus(ctx context.Context) []domain.MigrationStatusResponse {
	resp := []domain.MigrationStatusResponse{}

	for _, repo := range app.repos {
		resp = append(resp, repo.MigrationStatus(ctx))
	}

	return resp
}

func (app *Application) Repositories() []domain.RepositoryContract {
	return app.repos
}
