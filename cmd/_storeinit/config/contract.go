package config

const EnvConfigPath = "INITDB_CONFIG_PATH"

type ManagerContract interface {
	Configuration() Configuration
	Close() error
}
