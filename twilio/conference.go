package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/cousine/go-twilio/helpers"
	"io"
)

const (
	CONFERENCE_ENDPOINT = "Conferences"
)

type Conference struct {
	Sid          string `json:"sid"`
	FriendlyName string `json:"friendly_name"`
	DateCreated  string `json:"date_created"`
	DateUpdated  string `json:"date_updated"`
	AccountSid   string `json:"account_sid"`
	Status       string `json:"status"`
	Uri          string `json:"uri"`

	client *Client
}

func NewConference(client *Client, accountSid string) *Conference {
	return &Conference{
		client:     client,
		AccountSid: accountSid,
	}
}

func conferenceUri(accountSid string, conferenceSid string) string {
	return fmt.Sprintf("%v/%v/%v/%v", ACCOUNT_ENDPOINT, accountSid, CONFERENCE_ENDPOINT, conferenceSid)
}

func (conference *Conference) Client() *Client {
	return conference.client
}

func (conference *Conference) SetClient(client *Client) {
	conference.client = client
}

func (conference *Conference) FromJson(rawJson io.ReadCloser) error {
	jsonDecoder := json.NewDecoder(rawJson)
	err := jsonDecoder.Decode(conference)

	return err
}

func (client *Client) GetConference(accountSid string, conferenceSid string) (conference *Conference, err error) {
	conference = NewConference(client, accountSid)

	uri := generateEndpointUrl(conferenceUri(accountSid, conferenceSid))

	req, err := client.generateRequest("GET", uri, map[string]string{})
	if err != nil {
		return
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	if helpers.ResponseHasError(resp) {
		return nil, NewTwilioError(resp.Body)
	}

	err = conference.FromJson(resp.Body)
	if err != nil {
		return
	}

	return
}
