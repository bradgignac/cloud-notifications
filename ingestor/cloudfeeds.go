package ingestor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bradgignac/rcnotify/notifier"
)

// CloudFeeds ingests activity from Rackspace Cloud Feeds.
type CloudFeeds struct {
	Notifier notifier.Notifier
	Interval time.Duration
	User     string
	Key      string
	tenant   string
	token    string
	marker   string
}

// Start begins polling Cloud Feeds.
func (i *CloudFeeds) Start() error {
	err := i.authenticate()
	if err != nil {
		return err
	}

	for {
		i.readEvents()
		time.Sleep(i.Interval)
	}
}

func (i *CloudFeeds) authenticate() error {
	body := fmt.Sprintf(`{
		"auth": {
			"RAX-KSKEY:apiKeyCredentials": {
				"username": "%s",
				"apiKey": "%s"
			}
		}
	}`, i.User, i.Key)
	reader := strings.NewReader(body)

	res, err := http.Post("https://identity.api.rackspacecloud.com/v2.0/tokens", "application/json", reader)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var data map[string]interface{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&data)
	if err != nil {
		return err
	}

	access := data["access"].(map[string]interface{})
	token := access["token"].(map[string]interface{})
	tenant := token["tenant"].(map[string]interface{})

	i.tenant = tenant["id"].(string)
	i.token = token["id"].(string)

	log.Printf("Successfully authenticated user \"%s\", tenant \"%s\"", i.User, i.tenant)

	return nil
}

func (i *CloudFeeds) readEvents() {
	url := fmt.Sprintf("https://dfw.feeds.api.rackspacecloud.com/dbaas/events/%s/", i.tenant)
	if i.marker != "" {
		url = fmt.Sprintf("%s?marker=%s", url, i.marker)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create Cloud Feeds request - %v", err)
		return
	}

	req.Header.Set("X-Auth-Token", i.token)
	req.Header.Set("Accept", "application/vnd.rackspace.atom+json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to read Cloud Feeds events - %v", err)
		return
	}
	defer res.Body.Close()

	var data map[string]interface{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&data)
	if err != nil {
		log.Printf("Failed to parse Cloud Feeds response - %v", err)
		return
	}

	feed := data["feed"].(map[string]interface{})
	entries := feed["entry"].([]interface{})

	for _, v := range entries {
		entry := v.(map[string]interface{})
		content := entry["content"].(map[string]interface{})
		event := content["event"].(map[string]interface{})

		id := entry["id"]
		name := event["resourceName"]
		region := event["region"]
		action := event["rootAction"]

		switch action {
		case "trove.instance.delete":
			notification := fmt.Sprintf("%s was deleted from %s", name, region)
			i.Notifier.Notify(notification)
		}

		i.marker = id.(string)
	}

	log.Printf("Successfully polled %d event(s) for user \"%s\", tenant \"%s\", marker: \"%s\"", len(entries), i.User, i.tenant, i.marker)
}
