package notifier

import "log"

// Console logs notifications to the console.
type Console struct {
}

// Notify logs a notification to the console.
func (c *Console) Notify(n string) {
	log.Println(n)
}
