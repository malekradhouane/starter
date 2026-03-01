package pg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"storeinit/domain"

	"github.com/golang-migrate/migrate/v4"
)

func (adapter *Adapter) Migrate(ctx context.Context, req domain.MigrateRequest) domain.MigrateResponse {
	resp := domain.MigrateResponse{Repo: adapter}

	migration, err := migrate.New(adapter.migrationFolder, adapter.dbURI)
	if err != nil {
		resp.Err = err
		return resp
	}
	defer migration.Close()

	if ctx != nil {
		if t, ok := ctx.Deadline(); ok {
			migration.LockTimeout = time.Until(t)
		}
	}

	switch req.Action {
	case domain.ActionMigrateUp:
		if err := migration.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				resp.NoChange = true
				return resp
			}

			resp.Err = err
			return resp
		}

		return resp

	case domain.ActionMigrateForce:
		toVersion := int(adapter.schemaVersion)
		if err := migration.Force(toVersion); err != nil {
			resp.Err = err
			return resp
		}

		return resp

	case domain.ActionMigrateDown:
		if err := migration.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				resp.NoChange = true
				return resp
			}

			resp.Err = err
			return resp
		}

		return resp
	}

	resp.Err = fmt.Errorf("%w: unknown command: %s", ErrRepo, req.Action)

	return resp
}
