package config

const (
	DriverMongoDB    = "mongodb"
	DriverPostgreSQL = "pg"
)

type (
	Configuration struct {
		Storages []ConfStorages
	}

	ConfStorages struct {
		Name               string `yaml:"name"`
		Enabled            bool   `yaml:"enabled"`
		Driver             string `yaml:"driver"`
		DatabaseURI        string `yaml:"dbURI"`
		SchemaFolder       string `yaml:"schemaFolder"`
		ForceSchemaVersion uint   `yaml:"forceSchemaVersion"`
	}
)
