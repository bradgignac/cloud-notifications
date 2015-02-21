package notifier

// Notifier provides an interface for implementing multiple notification types.
type Notifier interface {
	Notify(n string)
}
