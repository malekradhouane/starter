package pg

import (
	"context"
	"errors"

	"storeinit/domain"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (adapter *Adapter) MigrationStatus(ctx context.Context) domain.MigrationStatusResponse {
	resp := domain.MigrationStatusResponse{Repo: adapter}

	migration, err := migrate.New(adapter.migrationFolder, adapter.dbURI)
	if err != nil {
		resp.Err = err
		return resp
	}
	defer migration.Close()

	currentVersion, dirty, err := migration.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			resp.SchemaVersion = -1
			return resp
		}

		resp.Err = err
		return resp
	}

	resp.SchemaVersion = int(currentVersion)
	resp.Dirty = dirty

	return resp
}
