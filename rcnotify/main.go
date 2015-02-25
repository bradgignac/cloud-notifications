package main

import (
	"log"
	"os"
	"time"

	"github.com/bradgignac/cloud-notifications/ingestor"
	"github.com/bradgignac/cloud-notifications/notifier"
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
			Name:  "rackspace-user",
			Usage: "Rackspace Cloud username",
		},
		cli.StringFlag{
			Name:  "rackspace-key",
			Usage: "Rackspace Cloud API key",
		},
		cli.StringFlag{
			Name:  "twilio-account",
			Usage: "Twilio account ID",
		},
		cli.StringFlag{
			Name:  "twilio-key",
			Usage: "Twilio API key",
		},
		cli.StringFlag{
			Name:  "twilio-from",
			Usage: "Twilio number",
		},
		cli.StringFlag{
			Name:  "twilio-to",
			Usage: "End-user number",
		},
	}

	app.Action = poll

	app.Run(os.Args)
}

func poll(c *cli.Context) {
	twAccount := arg(c, "twilio-account")
	twKey := arg(c, "twilio-key")
	twFrom := arg(c, "twilio-from")
	twTo := arg(c, "twilio-to")

	notifier := &notifier.Twilio{
		Account: twAccount,
		Token:   twKey,
		From:    twFrom,
		To:      twTo,
	}

	rsUser := arg(c, "rackspace-user")
	rsKey := arg(c, "rackspace-key")

	ingestor := &ingestor.CloudFeeds{
		Notifier: notifier,
		Interval: 10 * time.Second,
		User:     rsUser,
		Key:      rsKey,
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
