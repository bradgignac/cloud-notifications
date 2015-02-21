package main

import (
	"log"
	"os"
	"time"

	"github.com/bradgignac/rcnotify/ingestor"
	"github.com/bradgignac/rcnotify/notifier"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rcnotify"
	app.Usage = "Push notifications for Rackspace Cloud"
	app.Version = "0.0.0"
	app.Author = "Brad Gignac"
	app.Email = "bgignac@bradgignac.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "Rackspace Cloud username",
			EnvVar: "RCNOTIFY_USER",
		},
		cli.StringFlag{
			Name:   "key, k",
			Usage:  "Rackspace Cloud API key",
			EnvVar: "RCNOTIFY_KEY",
		},
	}

	app.Action = poll

	app.Run(os.Args)
}

func poll(c *cli.Context) {
	user := arg(c, "user")
	key := arg(c, "key")

	notifier := &notifier.Console{}
	ingestor := &ingestor.CloudFeeds{
		Notifier: notifier,
		Interval: 10 * time.Second,
		User:     user,
		Key:      key,
	}

	err := ingestor.Start()
	if err != nil {
		log.Fatalf("Failed to start ingestor: %v", err)
	}
}

func arg(c *cli.Context, name string) string {
	val := c.String(name)
	if val == "" {
		log.Fatalf("Parameter \"%s\" was not provided\n", name)
	}

	return val
}
