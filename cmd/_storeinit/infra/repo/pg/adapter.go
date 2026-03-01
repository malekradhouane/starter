package pg

import (
	"context"

	"storeinit/domain"
)

var _ domain.RepositoryContract = (*Adapter)(nil)

type Adapter struct {
	name            string
	dbURI           string
	migrationFolder string
	schemaVersion   uint
}

func (adapter *Adapter) Close(_ context.Context) error {
	return nil
}

func (adapter *Adapter) Driver() string {
	return "PostgreSQL"
}

func (adapter *Adapter) Name() string {
	return adapter.name
}
