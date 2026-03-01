package cli

import (
	"context"

	"storeinit/domain"
)

// What handler need from app

type ApplicationDeps interface {
	Migrate(ctx context.Context, req domain.MigrateRequest) []domain.MigrateResponse
	MigrationStatus(ctx context.Context) []domain.MigrationStatusResponse
	Repositories() []domain.RepositoryContract
}
