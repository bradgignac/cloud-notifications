package notifier

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Twilio sends notifications over SMS.
type Twilio struct {
	Account string
	Token   string
	From    string
	To      string
}

// Notify logs a notification to the console.
func (t *Twilio) Notify(n string) {
	uri := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.Account)
	form := url.Values{"To": {t.To}, "From": {t.From}, "Body": {n}}

	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		log.Printf("Failed to create Twilio request - %v", err)
		return
	}
	req.SetBasicAuth(t.Account, t.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to perform Twilio request - %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		log.Printf("Failed to perform Twilio request - %v", err)
	}

	var data map[string]interface{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&data)
	if err != nil {
		log.Printf("Failed to parse Twilio response - %v", err)
		return
	}

	log.Printf("Successfully sent Twilio notification with SID \"%s\"", data["sid"])
}
