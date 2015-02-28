package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// FactoryConfig represents type-specific configuration data used by different
// types of ingestors, notifiers, and coordinators.
type FactoryConfig struct {
	Type    string
	Options map[string]interface{}
}

// Config defines configuration data for Cloud Notifications.
type Config struct {
	Endpoints []string
	Feeds     []string
	Ingestor  FactoryConfig
	Notifier  FactoryConfig
}

// LoadYAML returns a Config struct from a YAML file.
func LoadYAML(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
