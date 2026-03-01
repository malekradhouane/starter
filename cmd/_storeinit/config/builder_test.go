package config

import (
	"log"
	"reflect"
	"testing"
)

func TestDefaultManager(t *testing.T) {
	testConfig := &Configuration{
		Storages: []ConfStorages{
			{
				Name:               "test",
				Enabled:            false,
				Driver:             "mongo",
				DatabaseURI:        "db://login@pass:localhost:12345/db",
				SchemaFolder:       "file://migrations",
				ForceSchemaVersion: 10,
			},
		},
	}

	cman, err := DefaultManager().
		WithTestConfiguration(testConfig).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	config := cman.Configuration()
	if !reflect.DeepEqual(*testConfig, config) {
		log.Fatalf("configurations must be the equal: %#v != %#v", *testConfig, config)
	}
}
