package config

import (
	"fmt"
	"os"
)

// Option represents a config option read from an unstructured JSON blob.
type Option struct {
	Key string
	Env string
}

// ReadOption retrieves a config option from an unstructured JSON blob.
func ReadOption(option Option, config map[string]interface{}) (string, error) {
	if val := os.Getenv(option.Env); val != "" {
		return val, nil
	} else if val, ok := config[option.Key]; ok {
		return val.(string), nil
	}

	return "", fmt.Errorf("Missing option \"%s\"", option.Key)
}

// ReadOptions retrieves a set of config options from an unstructured JSON blob.
func ReadOptions(options []Option, config map[string]interface{}) (map[string]string, error) {
	values := make(map[string]string, len(options))

	for _, opt := range options {
		val, err := ReadOption(opt, config)
		if err != nil {
			return nil, err
		}

		values[opt.Key] = val
	}

	return values, nil
}
