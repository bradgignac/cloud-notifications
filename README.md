# rcnotify

rcnotify is a simple command-line application that polls Rackspace Cloud Feeds for activity on your account and sends notifications based on events in the feed.

## Status

This project is pretty much a hack. To turn it into something more usable, the following features are needed:

- Tests.
- Code that isn't awful.
- Config syntax for specifying notifiers.
- Config syntax for specifying notification types.
- Support for more event feeds and more event types.
- Cooperative work across multiple worker threads and worker processes.

## Installation

To install Rackspace Cloud Notifications, you will need to have Go v1.3 or higher installed. Simply run:

```
$ go install github.com/bradgignac/rcnotify
```

`rcnotify` is now available in the `bin` directory of your `$GOPATH`.

## Usage

Start the daemon:

```
$ rcnotify -u MY_USER -k MY_API_KEY -p "my phone number"
```

## License

rcnotify is released under the [MIT License](LICENSE).
