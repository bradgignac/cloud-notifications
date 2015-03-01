package ingestor

import "fmt"

// Ingestor provides an interface for implementing multiple ingestor types.
type Ingestor interface {
	Start() error
}

// New creates an ingestor for the provided type and options.
func New(factoryType string, options map[string]interface{}) (Ingestor, error) {
	switch factoryType {
	case "rackspace":
		return NewRackspaceIngestor(options)
	}

	return nil, fmt.Errorf("No ingestor for type \"%s\"", factoryType)
}
