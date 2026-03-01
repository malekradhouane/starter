package config

type manager struct {
	conf       *Configuration
	configPath string
}

func (manager *manager) Configuration() Configuration {
	return *manager.conf
}

func (manager *manager) Close() error {
	manager.conf = nil

	return nil
}
