package ingestor

// Ingestor provides an interface for implementing multiple ingestor types.
type Ingestor interface {
	Start() error
}
