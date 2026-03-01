package app

import (
	"context"

	"storeinit/domain"
)

type ApplicationContract interface {
	Migrate(ctx context.Context, req domain.MigrateRequest) []domain.MigrateResponse
	MigrationStatus(ctx context.Context) []domain.MigrationStatusResponse
	Repositories() []domain.RepositoryContract
}
