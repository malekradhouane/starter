package configmanager

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"github.com/knadh/koanf/v2"
)

type ManagerWithKoanf struct {
	mu             sync.Mutex
	logger         ManagerLogger
	trippyConfig   *Trippy
	configRootDir  string
	environment    string
	configProvider *koanf.Koanf
}

func (manager *ManagerWithKoanf) WithLogger(logger ManagerLogger) *ManagerWithKoanf {
	manager.logger = logger
	return manager
}

func (manager *ManagerWithKoanf) WithTrippyConfig(trippyConfig *Trippy) *ManagerWithKoanf {
	manager.trippyConfig = trippyConfig
	return manager
}

func (manager *ManagerWithKoanf) WithConfigRoot(configRoot string) *ManagerWithKoanf {
	manager.configRootDir = configRoot
	return manager
}

func (manager *ManagerWithKoanf) WithEnvironment(env string) *ManagerWithKoanf {
	manager.environment = env
	return manager
}

func (manager *ManagerWithKoanf) Build() (*ManagerWithKoanf, error) {
	if manager.logger == nil {
		manager.logger = slog.Default()
	}

	if manager.trippyConfig != nil {
		return manager, nil
	}

	if manager.configRootDir == "" {
		if wd, err := os.Getwd(); err != nil {
			return nil, errors.Join(ErrConfigManager, err)
		} else {
			manager.configRootDir = wd
		}
	}

	if abs, err := filepath.Abs(manager.configRootDir); err != nil {
		return nil, errors.Join(ErrConfigManager, err)
	} else {
		manager.configRootDir = abs
	}

	if err := manager.loadAll(); err != nil {
		return nil, err
	}

	if manager.trippyConfig.BaseURL == "" {
		manager.trippyConfig.BaseURL = "http://localhost:8080"
	}

	return manager, nil
}

func DefaultManagerWithKonf() *ManagerWithKoanf {
	return &ManagerWithKoanf{
		logger:         slog.Default(),
		configProvider: koanf.New("."),
	}
}

func (cman *ManagerWithKoanf) Trippy() *Trippy {
	return cman.trippyConfig
}

func (cman *ManagerWithKoanf) Environment() string {
	return cman.environment
}

func (cman *ManagerWithKoanf) CustomerParam(key string) (any, bool) {
	val, found := cman.trippyConfig.Customer[key]
	return val, found
}
