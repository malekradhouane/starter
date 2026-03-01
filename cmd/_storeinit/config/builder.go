package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const DefaultConfigPath = "config.yaml"

func DefaultManager() *manager {
	return &manager{
		configPath: DefaultConfigPath,
	}
}

func (manager *manager) WithTestConfiguration(conf *Configuration) *manager {
	manager.conf = conf
	return manager
}

func (manager *manager) WithConfigPath(path string) *manager {
	manager.configPath = path
	return manager
}

func (manager *manager) Build() (*manager, error) {
	if manager.conf != nil {
		return manager, nil
	}

	if manager.configPath == "" {
		manager.configPath = "./config.yaml"
	}

	abs, err := filepath.Abs(manager.configPath)
	if err != nil {
		return nil, errors.Join(ErrConfig, err)
	}

	fileIO, err := os.Open(abs)
	if err != nil {
		return nil, errors.Join(ErrConfig, err)
	}
	defer fileIO.Close()

	// Load and decode configuration file.

	var conf Configuration
	switch ext := strings.ToLower(filepath.Ext(abs)); ext {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(fileIO).Decode(&conf)
	case ".json":
		err = json.NewDecoder(fileIO).Decode(&conf)
	default:
		err = fmt.Errorf("no decoder found for extension: %s", ext)
	}
	if err != nil {
		return nil, errors.Join(ErrConfig, err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_LOGIN")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host != "" && user != "" && password != "" && dbname != "" {
		if port == "" {
			port = "5432"
		}
		conf.Storages[0].DatabaseURI = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	}

	manager.conf = &conf

	return manager, nil
}
