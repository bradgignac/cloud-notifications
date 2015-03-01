# cloud-notifications

Cloud Notifications is a simple command-line application that polls Rackspace Cloud Feeds for activity on your account and sends notifications based on events in the feed. This currently notifies about deletes from Rackspace Cloud Databases.

## Status

This project is pretty much a hack. To turn it into something more usable, the following features are needed:

- Tests and benchmarks.
- Code that isn't awful.
- Specify start time on first poll.
- Config syntax for specifying notifiers.
- Config syntax for specifying notification types.
- Support for more event feeds and more event types.
- Cooperative work across multiple worker threads and worker processes.
- Respond to backpressure from a feed.
- Split ingestion from notification.
- Metrics exposed through expvar.

## Installation

To install Cloud Notifications, you will need to have Go v1.4 or higher installed. Simply run:

```
$ go install github.com/bradgignac/rcnotify
```

`rcnotify` is now available in the `bin` directory of your `$GOPATH`.

## Usage

Start the application:

```
$ rcnotify --rackspace-user RACKSPACE_USER \
    --rackspace-key RACKSPACE_KEY \
    --twilio-user TWILIO_ACCOUNT \
    --twilio-key TWILIO_TOKEN \
    --twilio-from TWILIO_NUMBER \
    --twilio-to MY_PHONE_NUMBER
```

## License

Cloud Notifications is released under the [MIT License](LICENSE).
