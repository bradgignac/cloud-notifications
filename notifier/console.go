package notifier

import "log"

// Console logs notifications to the console.
type Console struct {
}

// NewConsoleNotifier creates a new Console notifier with the provided options.
func NewConsoleNotifier(options map[string]interface{}) (*Console, error) {
	return &Console{}, nil
}

// Notify logs a notification to the console.
func (c *Console) Notify(n string) {
	log.Println(n)
}
