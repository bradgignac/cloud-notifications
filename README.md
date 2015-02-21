# rackspace-cloud-notifications

Rackspace Cloud Notifications is a simple daemon that polls Rackspace Cloud Feeds for activity on your account and sends alerts based on events in the feed.

## Installation

To install Rackspace Cloud Notifications, you will need to have Go v1.3 or higher installed. Simply run:

```
$ go install github.com/bradgignac/rackspace-cloud-notifications
```

`rcnotify`, the notification daemon, is now available in your `$GOPATH`.

## Usage

Start the daemon:

```
$ rcnotify -u MY_USER -k MY_API_KEY
```

## License

Slingshot is released under the [MIT License](LICENSE).
