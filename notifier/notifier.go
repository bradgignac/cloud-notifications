package notifier

import "fmt"

var notifiers = make(map[string]Notifier)

// Notifier provides an interface for implementing multiple notification types.
type Notifier interface {
	Notify(n string)
}

// New creates a notifier for the provided type and options.
func New(factoryType string, options map[string]interface{}) (Notifier, error) {
	switch factoryType {
	case "console":
		return NewConsoleNotifier(options)
	case "twilio":
		return NewTwilioNotifier(options)
	}

	return nil, fmt.Errorf("No notifier for type \"%s\"", factoryType)
}
