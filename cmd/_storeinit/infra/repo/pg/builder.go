package pg

import (
	"errors"
	"net/url"
)

// TODO:SB Use a connexion manager

type builder struct {
	Adapter
}

func NewAdapterBuilder() *builder { return &builder{} }

func (builder *builder) Build() (*Adapter, error) {
	// Validate and cleanup URI

	u, err := url.Parse(builder.dbURI)
	if err != nil {
		return nil, errors.Join(ErrRepo, err)
	}
	builder.dbURI = u.String()

	return &builder.Adapter, nil
}

func (builder *builder) SetName(name string) *builder {
	builder.name = name

	return builder
}

func (builder *builder) SetDatabaseURI(uri string) *builder {
	builder.dbURI = uri

	return builder
}

func (builder *builder) SetSchemaFolder(path string) *builder {
	builder.migrationFolder = path

	return builder
}

func (builder *builder) SetForceSchemaVersion(version uint) *builder {
	builder.schemaVersion = version

	return builder
}
